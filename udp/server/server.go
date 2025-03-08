package udpserver

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
)

func Run(port string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	socket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP address: %w", err)
	}
	defer socket.Close()

	buffer := make([]byte, internal.BufferSize)
	fmt.Println("Server running...")

	for {
		n, clientAddr, err := socket.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("failed to read from UDP: %v\n", err)
			continue
		}

		response := internal.ReverseWords(buffer[:n])
		fmt.Println(string(response))
		if _, err := socket.WriteToUDP(response, clientAddr); err != nil {
			fmt.Printf("failed to write to UDP: %v\n", err)
			continue
		}
	}
}
