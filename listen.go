package syncServer

import (
	"fmt"
	"net"

	"syncServer/message"
)

type listen struct {
	net.Listener
	*net.UDPConn
	addr net.Addr
}

func (l *listen) server(transport *transport) error {
	fmt.Printf("start listen with %s on %s.\n", l.addr.Network(), l.addr.String())

	switch l.addr.Network() {
	case "tcp":
		for {
			conn, err := l.Listener.Accept()
			if err != nil {
				return err
			}

			fmt.Printf("start accept with %s.\n", conn.RemoteAddr())
			transport.ctx.Send(trans, &message.Conn{Conn: conn})
			go transport.readTCP(conn)
		}
	case "udp":
		go transport.readUDP(l.UDPConn)
	}

	return nil
}

func (l *listen) close() {
	l.Close()
}

func NewListen(addr net.Addr) (*listen, error) {
	var l net.Listener
	var c *net.UDPConn
	var err error

	switch addr.Network() {
	case "tcp":
		laddr, _ := net.ResolveTCPAddr(addr.Network(), addr.String())
		l, err = net.ListenTCP(addr.Network(), laddr)
		if err != nil {
			return nil, err
		}
	case "udp":
		laddr, _ := net.ResolveUDPAddr(addr.Network(), addr.String())
		c, err = net.ListenUDP(addr.Network(), laddr)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(addr.String(), addr.Network())

	return &listen{addr: addr, UDPConn: c, Listener: l}, nil
}
