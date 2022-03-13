package example

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

const (
	ADDR = "127.0.0.1:8000"
)

func Test_syncServer(t *testing.T) {
	var tcpConn []net.Conn
	for i := 0; i < 2; i++ {
		conn, err := net.Dial("tcp", ADDR)
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(fmt.Sprintf("client%d", i)))
		tcpConn = append(tcpConn, conn)
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			msg := fmt.Sprintf("[client%d] message%d", i%2, i)
			buff := new(bytes.Buffer)
			buff.Write([]byte(msg))

			_, err := tcpConn[i%2].Write(buff.Bytes())
			if err != nil {
				panic(err)
			}

			pack := make([]byte, 1024)
			io.ReadFull(tcpConn[i%2], pack)
			fmt.Println("clinet", i%2, string(pack))
		}(i)
	}

	time.Sleep(time.Second*5)
}

func Benchmark_syncServer(b *testing.B) {
	var tcpConn []net.Conn
	for i := 0; i < 1000; i++ {
		conn, err := net.Dial("tcp", ADDR)
		if err != nil {
			panic(err)
		}
		tcpConn = append(tcpConn, conn)
	}

	for i := 0; i < b.N; i++ {
		msg := fmt.Sprintf("message%d", i)
		buff := new(bytes.Buffer)
		buff.Write([]byte(msg))

		_, err := tcpConn[i%1000].Write(buff.Bytes())
		if err != nil {
			panic(err)
		}

		pack := make([]byte, 1024)
		io.ReadFull(tcpConn[i%1000], pack)
	}
}