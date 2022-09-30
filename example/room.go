package main

import (
	"fmt"

	"syncServer"
	"syncServer/message"
	"syncServer/message/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type RoomSvc struct {
	actor.Context
	sync  *syncServer.Context
	userOpration map[uint32][]*pb.SyncMsg
}

func (r *RoomSvc) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("RoomSvc start succ!")
		r.Context = ctx

	case *message.Req:
		r.Request(msg)

	default:
		fmt.Printf("undefined message: %v\n", msg)
	}
}

func (r *RoomSvc) Request(req *message.Req) {
	switch msg := req.Content.(type) {
	case *pb.SyncMsg:
		r.Send(r.sync.Pid(), msg)
	}
}

func (r *RoomSvc) SvcName() string {
	return "RoomSvc"
}

func (r *RoomSvc) Pid() *actor.PID {
	return r.Context.Self()
}

func NewRoomSvc(ctx *syncServer.Context) syncServer.ServiceItf {
	return &RoomSvc{sync: ctx, userOpration: make(map[uint32][]*pb.SyncMsg)}
}
