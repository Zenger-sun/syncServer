package syncServer

import (
	"fmt"
	"net"

	"syncServer/message"

	"github.com/AsynkronIT/protoactor-go/actor"
)

var trans *actor.PID

type transport struct {
	ctx *actor.RootContext
	conn map[string]net.Conn
}

func (t *transport) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("sync transport start!")

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
		err := t.forward(msg)
		if err != nil {
			fmt.Println(err)
		}

	case *message.Res:
		err := t.write(msg)
		if err != nil {
			fmt.Println(err)
		}

	default:
		fmt.Printf("undefined message: %v\n", msg)
	}
}

func (t *transport) Pid() *actor.PID {
	return trans
}

func (t *transport) readTCP(conn net.Conn) {
	buff := make([]byte, 102400)

	defer func() {
		t.ctx.Send(t.Pid(), &message.Close{Addr: conn.RemoteAddr().String()})
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

		head := message.UnpackHead(buff)
		head.Addr = conn.RemoteAddr().String()
		req, err := message.UnpackReq(head, buff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		t.ctx.Send(t.Pid(), req)
	}
}

func (t *transport) readUDP(conn *net.UDPConn) {
	buff := make([]byte, 102400)

	for {
		_, remoteAddr, err := conn.ReadFromUDP(buff)
		switch err {
		case nil:
		default:
			t.ctx.Send(t.Pid(), &message.Close{Addr: remoteAddr.String()})
			fmt.Println(err)
			return
		}

		head := message.UnpackHead(buff)
		head.Addr = remoteAddr.String()
		req, err := message.UnpackReq(head, buff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		t.ctx.Send(t.Pid(), req)
	}
}

func (t *transport) forward(msg *message.Req) error {
	var err error

	switch msg.Head.WriteType {
	case message.SERVER_REQ:
		t.ctx.Send(svcM, msg)
	default:
		t.write(&message.Res{Head:msg.Head, Content: msg.Content})
	}

	return err
}

func (t *transport) write(msg *message.Res) error {
	var err error

	switch msg.Head.WriteType {
	case message.BROADCAST_ALL:
		return t.broadcast(msg)
	case message.BROADCAST_SINGLE:
		return t.broadcastSingle(msg)
	}

	return err
}

func (t *transport) broadcast(msg *message.Res) error {
	msg.Head.WriteType = message.BROADCAST_RES
	res := message.PackMsg(msg.Head, msg.Content)

	for _, conn := range t.conn {
		if _, err := conn.Write(res.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

func (t *transport) broadcastSingle(msg *message.Res) error {
	msg.Head.WriteType = message.BROADCAST_RES
	res := message.PackMsg(msg.Head, msg.Content)

	if conn, ok := t.conn[msg.Head.Addr]; ok {
		if _, err := conn.Write(res.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

func NewTransport() *transport {
	transport := &transport{conn: make(map[string]net.Conn)}
	transport.ctx = actor.NewActorSystem().Root
	trans = transport.ctx.Spawn(actor.PropsFromProducer(func() actor.Actor { return transport }))
	return transport
}
