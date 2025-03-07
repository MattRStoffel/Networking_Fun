package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":4201")
	if err != nil {
		fmt.Println("unable to resolve address")
		os.Exit(1)
	}
	socket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Cant hear you")
		os.Exit(1)
	}

	message := make([]byte, 1024)
	fmt.Println("Server running....")
	for {
		n, clientAddr, _ := socket.ReadFromUDP(message)
		str := string(message[:n])
		str = strings.ToUpper(str)
		socket.WriteToUDP([]byte(str), clientAddr)
	}
}
