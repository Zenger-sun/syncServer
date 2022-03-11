package syncServer

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"net"
)

type Context struct {
	*transport
	listen *listen
}

func (ctx *Context) Setup(addr *net.TCPAddr) {
	listen, err := NewListen(addr)
	if err != nil {
		panic(err)
	}

	ctx.listen = listen

	go listen.server(ctx.transport)
}

func (ctx *Context) Shutdown() {
	fmt.Println("syncServer Shutdown...")
	ctx.listen.Close()
}

func NewContext() *Context {
	transport := &transport{tcpConn: make(map[string]net.Conn)}
	transport.RootContext = actor.NewActorSystem().Root
	transport.Pid = transport.Spawn(actor.PropsFromProducer(func() actor.Actor { return transport }))

	return &Context{transport:transport}
}

