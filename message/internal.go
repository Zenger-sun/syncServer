package message

import (
	"net"

	"github.com/gogo/protobuf/proto"
)

type Req struct {
	Head    *Head
	Content proto.Message
}

type Res struct {
	Head    *Head
	Content proto.Message
}

type Conn struct {
	Conn net.Conn
}

type Close struct {
	Addr string
}
