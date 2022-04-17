package syncServer

import "net"

type cluster struct {
	listen *listen
}

func NewCluster(addr net.Addr) (*cluster, error) {
	l, err := NewListen(addr)
	if err != nil {
		return nil, err
	}

	return &cluster{listen: l}, nil
}
