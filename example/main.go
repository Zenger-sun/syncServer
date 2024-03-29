package main

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"net"
	"os"
	"os/signal"
	"syncServer/message"
	"syncServer/message/pb"
	"syscall"

	"syncServer"
)

const (
	USER_ID_NOT_SET = 1
	USER_ID_START   = 1000
)

type LoginSvc struct {
	actor.Context
	sync  *syncServer.Context
	users map[uint32]bool
}

func (l *LoginSvc) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("LoginSvc start succ!")
		l.Context = ctx

	case *message.Req:
		l.Request(msg)

	default:
		fmt.Printf("undefined message: %v\n", msg)
	}
}

func (l *LoginSvc) Request(req *message.Req) {
	switch msg := req.Content.(type) {
	case *pb.LoginReq:
		res := l.Login(msg)
		res.Head.Addr = req.Head.Addr

		l.Send(l.sync.Pid(), res)
	}
}

func (l *LoginSvc) Login(req *pb.LoginReq) *message.Res {
	var loginRes pb.LoginRes
	loginRes.Result = true

	fmt.Printf("Login req userId: %v\n", req.UserId)

	if req.UserId == USER_ID_NOT_SET {
		loginRes.UserId = uint32(len(l.users) + USER_ID_START)
		l.users[loginRes.UserId] = true
	} else {
		if _, ok := l.users[req.UserId]; !ok {
			loginRes.UserId = uint32(len(l.users) + USER_ID_START)
			l.users[loginRes.UserId] = true
		} else {
			loginRes.UserId = req.UserId
		}
	}

	fmt.Printf("Login res userId: %v\n", loginRes.UserId)

	return &message.Res{
		Head: &message.Head{
			MsgType:   message.LOGIN_RES_MSG,
			WriteType: message.BROADCAST_SINGLE,
			LockStep:  false,
		},
		Content: &loginRes,
	}
}

func (l *LoginSvc) SvcName() string {
	return "LoginSvc"
}

func (l *LoginSvc) Pid() *actor.PID {
	return l.Context.Self()
}

func NewLoginSvc(ctx *syncServer.Context) syncServer.ServiceItf {
	return &LoginSvc{sync: ctx, users: make(map[uint32]bool)}
}

func main() {
	sync := syncServer.NewContext()
	addr, _ := net.ResolveTCPAddr("tcp", "192.168.191.1:8000")

	sync.Setup(addr)
	sync.RegisterSvc(NewLoginSvc(sync))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		switch <-exit {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			sync.Shutdown()
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}
