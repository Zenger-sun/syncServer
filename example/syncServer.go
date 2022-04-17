package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"syncServer"
)

func main() {
	ctx := syncServer.NewContext()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	//cluster, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9501")
	ctx.Setup(addr, nil)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		switch <-exit {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			ctx.Shutdown()
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}