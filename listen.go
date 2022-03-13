package syncServer

import (
	"net"
	"syncServer/pb"
)

type listen struct {
	addr net.Addr
	tcp  *net.TCPListener
}

func (l *listen) server(transport *transport) error {
	for {
		conn, err := l.tcp.Accept()
		if err != nil {
			return err
		}

		transport.Send(transport.Pid, &pb.Conn{Conn: conn})
		go transport.read(conn)
	}
}

func (l *listen) close() {
	l.tcp.Close()
}

func NewListen(addr *net.TCPAddr) (*listen, error) {
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &listen{addr: addr, tcp: l}, nil
}
