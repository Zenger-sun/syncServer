package message

import (
	"net"

	"github.com/gogo/protobuf/proto"
)

type Head struct {
	Len       uint32
	MsgType   MsgType
	WriteType WriteType
	LockStep  bool

	Addr string
}

type Req struct {
	Head    *Head
	Content proto.Message
}

type Conn struct {
	Conn net.Conn
}

type Close struct {
	Addr string
}
