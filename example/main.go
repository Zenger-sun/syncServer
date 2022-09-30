package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"syncServer"
)

func main() {
	sync := syncServer.NewContext()
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8000")

	sync.Setup(addr)
	sync.RegisterSvc(NewLoginSvc(sync))
	sync.RegisterSvc(NewRoomSvc(sync))

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
