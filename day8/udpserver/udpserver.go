package main

import (
	"fmt"
	"net"
	"strings"
)

// UDP server

func main() {

	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen UDP failed, err:", err)
		return
	}
	defer socket.Close()
	// no need to Accept. send/receive data straight away
	var data [1024]byte
	for {
		n, addr, err := socket.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed, err:", err)
			return
		}

		fmt.Println(data[:n])
		reply := strings.ToUpper(string(data[:n]))
		// send
		socket.WriteToUDP([]byte(reply), addr)
	}
}
