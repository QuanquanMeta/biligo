package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	proto "github.com/biligo/day8/protocol"
)

// client side

func main() {
	stickyPackage()
}

func do() {
	// 1. coonect to server side
	socket := "127.0.0.1:20000"
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		fmt.Printf("dial %s failed\n", socket)
		return
	}
	// fmt.Println("client started")
	// 2 communication to the server

	var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input: ")
		msg, err = reader.ReadString('\n')

		if err == io.EOF {
			fmt.Println("end of file")
			return
		}

		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}

		conn.Write([]byte(msg))
	}

	conn.Close()
}

func stickyPackage() {
	socket := "127.0.0.1:20000"
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		fmt.Printf("dial %s failed\n", socket)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello, hell. How are you?`
		// using the protocol
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode failed, err:", err)
			return
		}
		conn.Write(b)
	}
}
