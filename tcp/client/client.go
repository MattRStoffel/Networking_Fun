package tcpclient

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
	"os"
)

func Run() {
	input := internal.GetInput()

	radder, err := net.ResolveTCPAddr("tcp", os.Getenv("ADDRESS")+":"+os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open socket
	socket, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := socket.Write([]byte(input)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)

	n, err := socket.Read(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(buffer[:n]))
}
