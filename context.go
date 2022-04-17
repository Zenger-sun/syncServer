package syncServer

import (
	"fmt"
	"net"
)

type Context struct {
	*transport
	listen *listen
	cluster *cluster
}

func (ctx *Context) Setup(addr net.Addr, cluster net.Addr) {
	// bulid server
	listen, err := NewListen(addr)
	if err != nil {
		panic(err)
	}
	ctx.listen = listen
	go listen.server(ctx.transport)

	// build cluster
	if cluster != nil {
		c, err := NewCluster(cluster)
		if err != nil {
			panic(err)
		}
		ctx.cluster = c
		go c.listen.server(ctx.transport)
	}
}

func (ctx *Context) Shutdown() {
	fmt.Println("syncServer Shutdown...")
	ctx.listen.close()

	if ctx.cluster != nil {
		ctx.cluster.listen.Addr()
	}

	fmt.Println("syncServer Closed!")
}

func NewContext() *Context {
	return &Context{transport: NewTransport()}
}
