package tcpserver

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
)

func Run(port string) error {
	raddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to resolve TCP address: %w", err)
	}

	listener, err := net.ListenTCP("tcp", raddr)
	if err != nil {
		return fmt.Errorf("failed to listen on TCP address: %w", err)
	}
	defer listener.Close()

	fmt.Println("Server running...")

	buffer := make([]byte, common.BufferSize)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept connection: %v\n", err)
			continue
		}

		go handleConnection(conn, buffer)
	}
}

func handleConnection(conn net.Conn, buffer []byte) {
	defer conn.Close()

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("failed to read from connection: %v\n", err)
		return
	}

	response := common.ReverseWords(buffer[:n])
	if _, err := conn.Write(response); err != nil {
		fmt.Printf("failed to write to connection: %v\n", err)
		return
	}
}
