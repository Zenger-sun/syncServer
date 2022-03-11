package example

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"testing"
)

const (
	ADDR = "127.0.0.1:8000"
)

func Test_syncServer(t *testing.T) {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("message%d", i)
		buff := new(bytes.Buffer)
		buff.Write([]byte(msg))

		_, err = conn.Write(buff.Bytes())
		if err != nil {
			panic(err)
		}

		pack := make([]byte, 1024)
		io.ReadFull(conn, pack)
		fmt.Println(string(pack))
	}
}

func Benchmark_syncServer(b *testing.B) {
	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		msg := fmt.Sprintf("message%d", i)
		buff := new(bytes.Buffer)
		buff.Write([]byte(msg))

		_, err = conn.Write(buff.Bytes())
		if err != nil {
			panic(err)
		}

		pack := make([]byte, 1024)
		io.ReadFull(conn, pack)
	}
}