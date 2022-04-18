package message

import (
	"syncServer/message/pb"

	"github.com/gogo/protobuf/proto"
)

func MsgRouter(message proto.Message) string {
	switch message.(type) {
	case *pb.LoginReq:
		return "LoginSvc"
	}

	return ""
}
