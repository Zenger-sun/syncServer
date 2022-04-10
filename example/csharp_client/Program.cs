using System;
using System.IO;
using System.Net.Sockets;
using System.Threading;
using Google.Protobuf;
using Goproto.Protoc.Proto3;

class syncClient
{
    private const string HOST_NAME = "127.0.0.1";
    private const int SERVER_PORT = 8000;
    private const int HEAD_LEN = 9;
    private const int PACK_SIZE = 1024;

    private static NetworkStream n;

    static void Main(string[] args)
    {
        TcpClient client = new TcpClient(HOST_NAME, SERVER_PORT);
        n = client.GetStream();

        new Thread(SendMsg).Start();
        ReceiveMsg();
    }

    static void ReceiveMsg()
    {
        byte[] pack = new byte[PACK_SIZE];

        while (true)
        {
            BinaryReader r = new BinaryReader(n);
            r.Read(pack);
            ReadMsg(pack);
            Array.Clear(pack, 0, 0);
        }
    }

    static void SendMsg()
    {
        BinaryWriter w = new BinaryWriter(n);

        while (true)
        {
            SyncMsg msg = new SyncMsg();
            msg.Content = "[client1]" + Console.ReadLine();

            w.Write(PackMsg(msg));
            w.Flush();
        }
    }

    static byte[] PackMsg(SyncMsg msg)
    {
        MemoryStream memoryStream = new MemoryStream();
        msg.WriteTo(memoryStream);

        MemoryStream result = new MemoryStream();
        result.Write(ConvHeadToBytes(msg.ToByteArray().Length, 1));
        result.Write(ConvHeadToBytes(1, 2));
        result.Write(ConvHeadToBytes(1,3));
        result.Write(ConvHeadToBytes(0,4));
        result.Write(msg.ToByteArray());

        return result.ToArray();
    }

    static byte[] ConvHeadToBytes(int n, int headType)
    {
        switch (headType)
        {
            case 1:
                var res1 = new byte[4];
                res1[0] = (byte)(n & 0xff);
                res1[1] = (byte)(n >> 8 & 0xff);
                res1[2] = (byte)(n >> 16 & 0xff);
                res1[3] = (byte)(n >> 24 & 0xff);
                return res1;
            case 2:
                var res2 = new byte[2];
                res2[0] = (byte)(n & 0xff);
                res2[1] = (byte)(n >> 8 & 0xff);
                return res2;
            case 3:
                var res3 = new byte[2];
                res3[0] = (byte)(n & 0xff);
                res3[1] = (byte)(n >> 8 & 0xff);
                return res3;
            case 4:
                var res4 = new byte[1];
                res4[0] = (byte)(n & 0xff);
                return res4;
        }

        return null;
    }

    static void ReadMsg(byte[] msg)
    {
        int msgLen = msg[0] | msg[1] << 8 | msg[2] << 16 | msg[3] << 24;
        
        byte[] content = msg[HEAD_LEN..(HEAD_LEN + msgLen)];

        SyncMsg syncMsg = new SyncMsg();
        syncMsg.MergeFrom(content);
        Console.WriteLine(syncMsg.Content);
    }
}

