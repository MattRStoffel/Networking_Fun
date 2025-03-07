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
		fmt.Println(err)
		os.Exit(1)
	}
	socket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	fmt.Println("Server running....")
	for {
		_, clientAddr, _ := socket.ReadFromUDP(buffer)
		if _, err := socket.WriteToUDP(internal.ReverseWords(buffer), clientAddr); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
