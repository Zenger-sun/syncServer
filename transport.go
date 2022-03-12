package syncServer

import (
	"fmt"
	"net"
	"syncServer/proto"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type transport struct {
	*actor.RootContext
	Pid     *actor.PID
	tcpConn []net.Conn
}

func (t *transport) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("sync start!")
	case *proto.Ack:
	case *proto.Req:
		fmt.Printf("Content: %s\n", msg.Content)
		for _, conn := range t.tcpConn {
			t.write(conn, []byte(msg.Content))
		}
	default:
		fmt.Printf("undefined msg: %v\n", msg)
	}
}

func (t *transport) read(conn net.Conn) {
	buff := make([]byte, 1024)
	defer conn.Close()

	for {
		_, err := conn.Read(buff)
		switch err {
		case nil:
		default:
			fmt.Println(err)
			return
		}

		t.Send(t.Pid, &proto.Req{Content: string(buff)})
	}
}

func (t *transport) write(conn net.Conn, msg []byte) {
	conn.Write(msg)
}
