package message

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/golang/protobuf/proto"
)

const (
	PACK_MAX_LEN = 102400
	HEAD_LEN = 9
)

func UnpackHead(msg []byte) (*Head, error) {
	var head Head

	head.Len = uint32(msg[0]) | uint32(msg[1])<<8 | uint32(msg[2])<<16 | uint32(msg[3])<<24
	head.MsgType = MsgType(uint16(msg[4]) | uint16(msg[5])<<8)
	head.WriteType = WriteType(uint16(msg[6]) | uint16(msg[7])<<8)
	if msg[8] == 0 {
		head.LockStep = false
	} else {
		head.LockStep = true
	}

	if head.Len > (PACK_MAX_LEN + HEAD_LEN) {
		return nil, errors.New("msg len error!")
	}
	if head.MsgType >= MSG_TYPE_END {
		return nil, errors.New("msg type error!")
	}
	if head.WriteType >= WRITE_TYPE_END {
		return nil, errors.New("write type error!")
	}

	return &head, nil
}

func UnpackReq(head *Head, msg []byte) (*Req, error) {
	req := &Req{}
	req.Head = head
	req.Content = head.MsgType.GetMsgStruct()

	err := proto.Unmarshal(msg[HEAD_LEN:HEAD_LEN+head.Len], req.Content)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func PackHead(head *Head) *bytes.Buffer {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.LittleEndian, head.Len)
	binary.Write(buff, binary.LittleEndian, head.MsgType)
	binary.Write(buff, binary.LittleEndian, head.WriteType)
	if head.LockStep {
		binary.Write(buff, binary.LittleEndian, uint8(1))
	} else {
		binary.Write(buff, binary.LittleEndian, uint8(0))
	}

	return buff
}

func PackMsg(head *Head, message proto.Message) *bytes.Buffer {
	msg, _ := proto.Marshal(message)
	head.Len = uint32(len(msg))

	buff := PackHead(head)
	buff.Write(msg)

	return buff
}
