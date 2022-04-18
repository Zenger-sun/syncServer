package syncServer

import (
	"fmt"
	"net"
)

type Context struct {
	*transport
	*serviceM
	listen *listen
}

func (ctx *Context) Setup(addr net.Addr) {
	listen, err := NewListen(addr)
	if err != nil {
		panic(err)
	}
	ctx.listen = listen
	go listen.server(ctx.transport)

	fmt.Println("sync started!")
}

func (ctx *Context) Shutdown() {
	fmt.Println("syncServer Shutdown...")
	ctx.listen.close()
	fmt.Println("syncServer Closed!")
}

func NewContext() *Context {
	return &Context{
		transport: NewTransport(),
		serviceM:   NewService(),
	}
}
