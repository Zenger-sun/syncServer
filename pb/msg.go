package pb

import "net"

type WriteType uint32

const (
	_ WriteType = iota
	BROADCAST_ALL
	BROADCAST_RANGE
	BROADCAST_SINGLE
)

type Head struct {
	Len       uint32
	WriteType WriteType
	LockStep  bool
}

type Req struct {
	Head    Head
	Content []byte
}

type Res struct {
	Head    Head
	Content []byte
}

type Conn struct {
	Conn net.Conn
}

type Close struct {
	Addr string
}
