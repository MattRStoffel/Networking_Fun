package udpserver

import (
	"fmt"
	"net"
	"os"
)

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func Run() {
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
		_, clientAddr, _ := socket.ReadFromUDP(message)
		reverseBytes(message)
		socket.WriteToUDP(message, clientAddr)
	}
}
