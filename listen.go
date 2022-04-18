package syncServer

import (
	"fmt"
	"net"

	"syncServer/message"
)

type listen struct {
	net.Listener
	addr net.Addr
}

func (l *listen) server(transport *transport) error {
	fmt.Printf("start listen with %s on %s.\n", l.Addr().Network(), l.Addr().String())

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		transport.ctx.Send(trans, &message.Conn{Conn: conn})
		go transport.read(conn)
	}
}

func (l *listen) close() {
	l.Close()
}

func NewListen(addr net.Addr) (*listen, error) {
	var l net.Listener
	var err error

	if addr.Network() == "tcp" {
		laddr, _ := net.ResolveTCPAddr("tcp", addr.String())
		l, err = net.ListenTCP("tcp", laddr)
		if err != nil {
			return nil, err
		}
	}

	return &listen{addr: addr, Listener: l}, nil
}
