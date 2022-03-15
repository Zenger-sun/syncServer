package syncServer

import (
	"fmt"
	"net"

	"syncServer/pb"

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

	case *pb.Conn:
		if _, ok := t.conn[msg.Conn.RemoteAddr().String()]; !ok {
			t.conn[msg.Conn.RemoteAddr().String()] = msg.Conn
		}

	case *pb.Close:
		if _, ok := t.conn[msg.Addr]; ok {
			t.conn[msg.Addr].Close()
		}
		delete(t.conn, msg.Addr)

	case *pb.Req:
		err := t.write(msg)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Printf("undefined msg: %v\n", msg)
	}
}

func (t *transport) read(conn net.Conn) {
	buff := make([]byte, 1024)

	defer func() {
		t.Send(t.Pid, &pb.Close{Addr: conn.RemoteAddr().String()})
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
		req, err := UnpackReq(head, buff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		t.Send(t.Pid, req)
	}
}

func (t *transport) write(msg *pb.Req) error {
	var err error

	switch msg.Head.WriteType {
	case pb.BROADCAST_ALL:
		err = t.broadcast(msg)
	case pb.SERVER_REQ:
		err = t.serverReq(msg)
	}

	return err
}

func (t *transport) broadcast(msg *pb.Req) error {
	msg.Head.WriteType = pb.BROADCAST_RES
	res := PackMsg(msg.Head, msg.Content)

	for _, conn := range t.conn {
		if _, err := conn.Write(res.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

func (t *transport) serverReq(msg *pb.Req) error {
	switch msg := msg.Content.(type) {
	case *pb.ServerReq:
		fmt.Println(msg.Content)
	}

	return nil
}

func NewTransport() *transport {
	transport := &transport{conn: make(map[string]net.Conn)}
	transport.RootContext = actor.NewActorSystem().Root
	transport.Pid = transport.Spawn(actor.PropsFromProducer(func() actor.Actor { return transport }))
	return transport
}
