package udpserver

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
	"os"
)

func Run() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+os.Getenv("PORT"))
	if err != nil {
		fmt.Println("unable to resolve address")
		os.Exit(1)
	}
	socket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Cant hear you")
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	fmt.Println("Server running....")
	for {
		_, clientAddr, _ := socket.ReadFromUDP(buffer)
		socket.WriteToUDP(internal.ReverseWords(buffer), clientAddr)
	}
}
