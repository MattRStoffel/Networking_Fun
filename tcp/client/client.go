package tcpclient

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
)

func Run(addr string, port string) error {
	input := common.GetInput()

	radder, err := net.ResolveTCPAddr("tcp", addr+":"+port)
	if err != nil {
		return fmt.Errorf("failed to resolve TCP address: %w", err)
	}

	socket, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		return fmt.Errorf("failed to dial TCP: %w", err)
	}
	defer socket.Close()

	if _, err := socket.Write([]byte(input)); err != nil {
		return fmt.Errorf("failed to write to socket: %w", err)
	}

	buffer := make([]byte, common.BufferSize)

	n, err := socket.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read from socket: %w", err)
	}

	fmt.Println(string(buffer[:n]))
	return nil
}
