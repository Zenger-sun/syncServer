package syncServer

import (
	"fmt"

	"syncServer/message"

	"github.com/AsynkronIT/protoactor-go/actor"
)

var svcM *actor.PID

type ServiceItf interface {
	Pid() *actor.PID
	Receive(ctx actor.Context)
	SvcName() string
}

type serviceM struct {
	ctx *actor.RootContext
	svc map[string]ServiceItf
}

func (s *serviceM) RegisterSvc(svc ServiceItf) error {
	s.ctx.Send(svcM, svc)
	return nil
}

func (s *serviceM) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("sync service start!")

	case ServiceItf:
		s.svc[msg.SvcName()] = msg
		ctx.Spawn(actor.PropsFromProducer(func() actor.Actor { return msg }))

	case *message.Req:
		svc := s.findSvc(message.MsgRouter(msg.Content))
		if svc != nil {
			ctx.Send(svc.Pid(), msg)
		}

	default:
		fmt.Printf("undefined message: %v\n", msg)
	}
}

func (s *serviceM) findSvc(name string) ServiceItf {
	if svc, ok := s.svc[name]; ok {
		return svc
	}

	return nil
}

func NewService() *serviceM {
	serviceM := &serviceM{svc: make(map[string]ServiceItf)}
	serviceM.ctx = actor.NewActorSystem().Root
	svcM = serviceM.ctx.Spawn(actor.PropsFromProducer(func() actor.Actor { return serviceM }))
	return serviceM
}
