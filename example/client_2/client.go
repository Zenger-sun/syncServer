package main

import (
	"bytes"
	"fmt"
	"net"
	"syncServer"
	"syncServer/message"
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
		fmt.Scanf("%s", &msg)
		if msg == "close" {
			conn.Close()
			continue
		}

		buff.WriteString("[client2] ")
		buff.WriteString(msg)

		req := &msg.SyncMsg{Content: buff.String()}
		head := &msg.Head{
			MsgType:   msg.SYNC_MSG,
			WriteType: msg.BROADCAST_ALL,
			LockStep:  false,
		}

		_, err := conn.Write(syncServer.PackMsg(head, req).Bytes())
		if err != nil {
			panic(err)
		}

		buff.Reset()
	}
}
