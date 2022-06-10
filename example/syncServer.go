package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"syncServer"
	"syncServer/message"
	"syncServer/message/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type LoginSvc struct {
	actor.Context
	sync  *syncServer.Context
	users map[uint32]bool
}

func (l *LoginSvc) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
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

func (l *LoginSvc) Login(req *pb.LoginReq) *message.Res  {
	var loginRes pb.LoginRes
	loginRes.Result = true

	if _, ok := l.users[req.UserId]; !ok {
		if req.UserId != 0 {
			loginRes.UserId = req.UserId
		} else {
			loginRes.UserId = uint32(len(l.users)+1)
		}

		l.users[loginRes.UserId] = true
	} else {
		loginRes.UserId = req.UserId
	}

	return &message.Res{
		Head:    &message.Head{
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
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")

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
