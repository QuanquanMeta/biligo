package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	proto "github.com/biligo/day8/protocol"
)

// tcp server

func processProtocolConn(conn net.Conn) {
	// 3. communicate to cliet
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {

		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode failed, err:", err)
			return
		}
		fmt.Printf("Received from client: %s\n", msg)
	}
}

func processConn(conn net.Conn) {
	// 3. communicate to cliet
	defer conn.Close()
	var tmp [128]byte
	for {
		n, err := conn.Read(tmp[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed")
			return
		}

		fmt.Printf("Received from client: %s\n", string(tmp[:n]))
	}

}

func main() {
	// 1. start service at a socket
	socket := "127.0.0.1:20000"
	listener, err := net.Listen("tcp", socket)
	if err != nil {
		fmt.Printf("start server on %s, failed\n", socket)
		return
	}
	fmt.Println("Server started")
	// 2. wait for others to connect
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed")
			continue
		}

		go processProtocolConn(conn)
	}
}
