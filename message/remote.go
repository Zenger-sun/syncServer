package message

import (
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
	LOGIN_MSG
)

func (m MsgType) GetMsgStruct() proto.Message {
	switch m {
	case SYNC_MSG:
		return &SyncMsg{}
	case LOGIN_MSG:
		return &LoginReq{}
	}

	return nil
}
