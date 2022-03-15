package syncServer

import (
	"bytes"
	"encoding/binary"
	"syncServer/pb"

	"github.com/golang/protobuf/proto"
)

const (
	HEAD_LEN = 9
)

func UnpackHead(msg []byte) *pb.Head {
	var head pb.Head

	head.Len = uint32(msg[0]) | uint32(msg[1])<<8 | uint32(msg[2])<<16 | uint32(msg[3])<<24
	head.MsgType = pb.MsgType(uint16(msg[4]) | uint16(msg[5])<<8)
	head.WriteType = pb.WriteType(uint16(msg[6]) | uint16(msg[7])<<8)
	if msg[8] == 0 {
		head.LockStep = false
	} else {
		head.LockStep = true
	}

	return &head
}

func UnpackReq(head *pb.Head, msg []byte) (*pb.Req, error) {
	req := &pb.Req{}
	req.Head = head
	req.Content = head.MsgType.GetMsgStruct()

	err := proto.Unmarshal(msg[HEAD_LEN:HEAD_LEN+head.Len], req.Content)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func PackHead(head *pb.Head) *bytes.Buffer {
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

func PackMsg(head *pb.Head, message proto.Message) *bytes.Buffer {
	msg, _ := proto.Marshal(message)
	head.Len = uint32(len(msg))

	buff := PackHead(head)
	buff.Write(msg)

	return buff
}
