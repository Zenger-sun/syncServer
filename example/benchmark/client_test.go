package benchmark

import (
	"net"
	"syncServer/message"
	"syncServer/message/pb"
	"testing"
)

const (
	ADDR = "127.0.0.1:8000"
)

func Benchmark_client(b *testing.B) {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		panic(err)
	}

	head := &message.Head{
		MsgType:   message.LOGIN_REQ_MSG,
		WriteType: message.SERVER_REQ,
		LockStep:  false,
	}

	req := &pb.LoginReq{UserId: 1}
	go func() {
		for {
			pack := make([]byte, 102400)
			_, err = conn.Read(pack)
			if err != nil {
				continue
			}

			head := message.UnpackHead(pack)
			_, err = message.UnpackReq(head, pack)
			if err != nil {
				continue
			}
		}
	}()

	for i:=0; i < b.N; i++ {
		_, err := conn.Write(message.PackMsg(head, req).Bytes())
		if err != nil {
			panic(err)
		}
	}
}
