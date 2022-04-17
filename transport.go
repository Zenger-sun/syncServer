package syncServer

import (
	"fmt"
	"net"

	"syncServer/message"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type transport struct {
	*actor.RootContext
	Pid  *actor.PID
	conn map[string]net.Conn
}

func (t *transport) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("sync start!")

	case *message.Conn:
		if _, ok := t.conn[msg.Conn.RemoteAddr().String()]; !ok {
			t.conn[msg.Conn.RemoteAddr().String()] = msg.Conn
		}

	case *message.Close:
		if _, ok := t.conn[msg.Addr]; ok {
			t.conn[msg.Addr].Close()
		}
		delete(t.conn, msg.Addr)

	case *message.Req:
		err := t.write(msg)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Printf("undefined message: %v\n", msg)
	}
}

func (t *transport) read(conn net.Conn) {
	buff := make([]byte, 1024)

	defer func() {
		t.Send(t.Pid, &message.Close{Addr: conn.RemoteAddr().String()})
		conn.Close()
	}()

	for {
		_, err := conn.Read(buff)
		switch err {
		case nil:
		default:
			fmt.Println(err)
			return
		}

		head := UnpackHead(buff)
		head.Addr = conn.RemoteAddr().String()
		req, err := UnpackReq(head, buff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		t.Send(t.Pid, req)
	}
}

func (t *transport) write(msg *message.Req) error {
	var err error

	switch msg.Head.WriteType {
	case message.BROADCAST_ALL:
		err = t.broadcast(msg)
	case message.SERVER_REQ:
		err = t.serverReq(msg)
	}

	return err
}

func (t *transport) broadcast(msg *message.Req) error {
	msg.Head.WriteType = message.BROADCAST_RES
	res := PackMsg(msg.Head, msg.Content)

	for _, conn := range t.conn {
		if _, err := conn.Write(res.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

func (t *transport) serverReq(msg *message.Req) error {
	switch msg := msg.Content.(type) {
	case *message.LoginReq:
		fmt.Println(msg.String())
	}

	return nil
}

func NewTransport() *transport {
	transport := &transport{conn: make(map[string]net.Conn)}
	transport.RootContext = actor.NewActorSystem().Root
	transport.Pid = transport.Spawn(actor.PropsFromProducer(func() actor.Actor { return transport }))
	return transport
}
