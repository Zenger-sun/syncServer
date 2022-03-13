package proto

import "net"

type Req struct {
	Content string
}

type Res struct {
	Content string
}

type Conn struct {
	Conn net.Conn
}

type Close struct {
	Addr string
}
