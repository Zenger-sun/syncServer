package main

import (
	"bytes"
	"fmt"
	"net"
	"syncServer"
	"syncServer/pb"
)

const (
	ADDR = "127.0.0.1:8000"
)

func main() {

	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		panic(err)
	}

	pack := make([]byte, 1024)
	go func() {
		for {
			_, err := conn.Read(pack)
			if err != nil {
				continue
			}

			head := syncServer.UnpackHead(pack)
			fmt.Printf("head: %+v\t", head)
			req, err := syncServer.UnpackReq(head, pack)
			if err != nil {
				continue
			}

			fmt.Println(req.Content)
		}
	}()

	msg := ""
	buff := new(bytes.Buffer)
	for {
		head := &pb.Head{
			MsgType:   pb.SYNC_MSG,
			WriteType: pb.BROADCAST_ALL,
			LockStep:  false,
		}

		fmt.Scanf("%s", &msg)
		if msg == "close" {
			conn.Close()
			continue
		} else if msg == "server" {
			head.MsgType = pb.SERVER_MSG
			head.WriteType = pb.SERVER_REQ
		}

		buff.WriteString("[client1] ")
		buff.WriteString(msg)

		req := &pb.SyncMsg{Content:buff.String()}

		_, err := conn.Write(syncServer.PackMsg(head, req).Bytes())
		if err != nil {
			panic(err)
		}

		buff.Reset()
	}
}
