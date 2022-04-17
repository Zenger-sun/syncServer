package main

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/proto"
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
		head := &message.Head{
			MsgType:   message.SYNC_MSG,
			WriteType: message.BROADCAST_ALL,
			LockStep:  false,
		}

		fmt.Scanf("%s", &msg)
		if msg == "close" {
			conn.Close()
			continue
		} else if msg == "login" {
			head.MsgType = message.LOGIN_MSG
			head.WriteType = message.SERVER_REQ
		}

		buff.WriteString("[client1] ")
		buff.WriteString(msg)

		var req proto.Message
		if msg == "login" {
			req = &message.LoginReq{UserId: 1}
		} else {
			req = &message.SyncMsg{Content: buff.String()}
		}

		_, err := conn.Write(syncServer.PackMsg(head, req).Bytes())
		if err != nil {
			panic(err)
		}

		buff.Reset()
	}
}
