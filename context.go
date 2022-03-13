package syncServer

import (
	"fmt"
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
	ctx.listen.close()
}

func NewContext() *Context {
	return &Context{transport: NewTransport()}
}
