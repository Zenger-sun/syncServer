package main

import (
	"bytes"
	"fmt"
	"net"

	"syncServer/message"
	"syncServer/message/pb"

	"github.com/gogo/protobuf/proto"
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

			head := message.UnpackHead(pack)
			fmt.Printf("head: %+v\t", head)
			req, err := message.UnpackReq(head, pack)
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
			MsgType:   message.TEST_MSG,
			WriteType: message.BROADCAST_ALL,
			LockStep:  false,
		}

		fmt.Scanf("%s", &msg)
		if msg == "close" {
			conn.Close()
			continue
		} else if msg == "login" {
			head.MsgType = message.LOGIN_REQ_MSG
			head.WriteType = message.SERVER_REQ
		}

		buff.WriteString("[client1] ")
		buff.WriteString(msg)

		var req proto.Message
		if msg == "login" {
			req = &pb.LoginReq{UserId: 1}
		} else {
			req = &pb.TestMsg{Content: buff.String()}
		}

		_, err := conn.Write(message.PackMsg(head, req).Bytes())
		if err != nil {
			panic(err)
		}

		buff.Reset()
	}
}
