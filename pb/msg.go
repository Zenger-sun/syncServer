package pb

import (
	"net"

	"github.com/golang/protobuf/proto"
)

type WriteType uint16

const (
	_ WriteType = iota

	// forward
	BROADCAST_ALL
	BROADCAST_RANGE
	BROADCAST_SINGLE
	BROADCAST_RES

	// server
	SERVER_REQ
)

type MsgType uint16

const (
	_ MsgType = iota
	SYNC_MSG
	SERVER_MSG
)

func (m MsgType) GetMsgStruct() proto.Message {
	switch m {
	case SYNC_MSG:
		return &SyncMsg{}
	case SERVER_MSG:
		return &ServerReq{}
	}

	return nil
}

// ******* remote *******

type Head struct {
	Len       uint32
	MsgType   MsgType
	WriteType WriteType
	LockStep  bool
}

type Req struct {
	Head    *Head
	Content proto.Message
}

// ******* internal *******

type Conn struct {
	Conn net.Conn
}

type Close struct {
	Addr string
}
