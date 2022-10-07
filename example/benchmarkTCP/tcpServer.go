package main

import (
	"fmt"
	"net"
	"syncServer/message"
	"syncServer/message/pb"
)

const ADDR = "127.0.0.1:8000"

var users map[uint32]bool

func main() {
	users = make(map[uint32]bool)
	laddr, _ := net.ResolveTCPAddr("tcp", ADDR)
	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}

		fmt.Printf("start accept with %s.\n", conn.RemoteAddr())
		go read(conn)
	}
}

func read(conn net.Conn) {
	buff := make([]byte, 102400)

	for {
		_, err := conn.Read(buff)
		switch err {
		case nil:
		default:
			fmt.Println(err)
			return
		}

		head, _ := message.UnpackHead(buff)
		head.Addr = conn.RemoteAddr().String()
		req, err := message.UnpackReq(head, buff)
		if err != nil {
			fmt.Println(err)
			continue
		}

		res := Login(req.Content.(*pb.LoginReq))
		res.Head.WriteType = message.BROADCAST_RES
		conn.Write(message.PackMsg(res.Head, res.Content).Bytes())
	}
}

func Login(req *pb.LoginReq) *message.Res {
	var loginRes pb.LoginRes
	loginRes.Result = true

	if req.UserId == 1 {
		loginRes.UserId = uint32(len(users) + 1000)
		users[loginRes.UserId] = true
	} else {
		if _, ok := users[req.UserId]; !ok {
			loginRes.UserId = uint32(len(users) + 1000)
			users[loginRes.UserId] = true
		} else {
			loginRes.UserId = req.UserId
		}
	}

	return &message.Res{
		Head: &message.Head{
			MsgType:   message.LOGIN_RES_MSG,
			WriteType: message.BROADCAST_SINGLE,
			LockStep:  false,
		},
		Content: &loginRes,
	}
}