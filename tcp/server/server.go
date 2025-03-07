package tcpserver

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
	"os"
)

func Run() {
	raddr, err := net.ResolveTCPAddr("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	socket, err := net.ListenTCP("tcp", raddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	fmt.Println("Server running...")
	for {
		caddr, err := socket.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if _, err := caddr.Read(buffer); err != nil {
			fmt.Println(err)
			continue
		}
		if _, err := caddr.Write(internal.ReverseWords(buffer)); err != nil {
			fmt.Println(err)
			continue
		}
		if err := caddr.Close(); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
