package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client

func main() {
	// dial

	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("connect to server failed, err:", err)
		return
	}
	defer socket.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pls Input:")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		// receive the data

		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("received reply msg failed, err:", err)
			return
		}
		fmt.Println("received msg: ", string(reply[:n]))

	}
}
