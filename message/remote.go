package message

import (
	"syncServer/message/pb"

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

	WRITE_TYPE_END
)

type MsgType uint16

const (
	_ MsgType = iota
	TEST_MSG
	SYNC_MSG
	LOGIN_REQ_MSG
	LOGIN_RES_MSG

	MSG_TYPE_END
)

func (m MsgType) GetMsgStruct() proto.Message {
	switch m {
	case TEST_MSG:
		return &pb.TestMsg{}
	case SYNC_MSG:
		return &pb.SyncMsg{}
	case LOGIN_REQ_MSG:
		return &pb.LoginReq{}
	case LOGIN_RES_MSG:
		return &pb.LoginRes{}
	}

	return nil
}

type Head struct {
	Len       uint32
	MsgType   MsgType
	WriteType WriteType
	LockStep  bool

	Addr string
}