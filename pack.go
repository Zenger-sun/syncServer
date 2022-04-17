package syncServer

import (
	"bytes"
	"encoding/binary"
	"syncServer/message"

	"github.com/golang/protobuf/proto"
)

const (
	HEAD_LEN = 9
)

func UnpackHead(msg []byte) *message.Head {
	var head message.Head

	head.Len = uint32(msg[0]) | uint32(msg[1])<<8 | uint32(msg[2])<<16 | uint32(msg[3])<<24
	head.MsgType = message.MsgType(uint16(msg[4]) | uint16(msg[5])<<8)
	head.WriteType = message.WriteType(uint16(msg[6]) | uint16(msg[7])<<8)
	if msg[8] == 0 {
		head.LockStep = false
	} else {
		head.LockStep = true
	}

	return &head
}

func UnpackReq(head *message.Head, msg []byte) (*message.Req, error) {
	req := &message.Req{}
	req.Head = head
	req.Content = head.MsgType.GetMsgStruct()

	err := proto.Unmarshal(msg[HEAD_LEN:HEAD_LEN+head.Len], req.Content)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func PackHead(head *message.Head) *bytes.Buffer {
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

func PackMsg(head *message.Head, message proto.Message) *bytes.Buffer {
	msg, _ := proto.Marshal(message)
	head.Len = uint32(len(msg))

	buff := PackHead(head)
	buff.Write(msg)

	return buff
}
