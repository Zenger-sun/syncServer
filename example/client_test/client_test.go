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
	for i := 0; i < 100; i++ {
		conn, err := net.Dial("tcp", ADDR)
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(fmt.Sprintf("client_test%d", i)))
		tcpConn = append(tcpConn, conn)
	}

	for i := 0; i < 1000; i++ {
		go func(i int) {
			msg := fmt.Sprintf("[client_test%d] message%d", i%2, i)
			buff := new(bytes.Buffer)
			buff.Write([]byte(msg))

			_, err := tcpConn[i%100].Write(buff.Bytes())
			if err != nil {
				panic(err)
			}

			pack := make([]byte, 1024)
			io.ReadFull(tcpConn[i%100], pack)
			fmt.Println("clinet", i%100, string(pack))
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
		conn.Write([]byte(fmt.Sprintf("client_test%d", i)))
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