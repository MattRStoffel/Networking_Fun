package udpclient

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
)

func Run(addr string, port string) error {
	input := common.GetInput()

	udpAddr, err := net.ResolveUDPAddr("udp", addr+":"+port)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	connection, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return fmt.Errorf("failed to dial UDP: %w", err)
	}
	defer connection.Close()

	if _, err := connection.Write([]byte(input)); err != nil {
		return fmt.Errorf("failed to write to UDP connection: %w", err)
	}

	buffer := make([]byte, common.BufferSize)
	n, _, err := connection.ReadFromUDP(buffer)
	if err != nil {
		return fmt.Errorf("failed to read from UDP connection: %w", err)
	}

	fmt.Println(string(buffer[:n]))
	return nil
}
