package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

const (
	ADDR = "127.0.0.1:8000"
)

func main() {
	msg := ""
	buff := new(bytes.Buffer)

	conn, err := net.Dial("tcp", ADDR)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			pack := make([]byte, 1024)
			io.ReadFull(conn, pack)
			fmt.Println(string(pack))
		}
	}()

	for {
		fmt.Scanf("%s", &msg)
		buff.WriteString("[client2] ")
		buff.WriteString(msg)

		_, err := conn.Write(buff.Bytes())
		if err != nil {
			panic(err)
		}

		buff.Reset()
	}
}
