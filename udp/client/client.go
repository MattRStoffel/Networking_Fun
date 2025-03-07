package udpclient

import (
	"NetworkingFun/internal"
	"fmt"
	"net"
	"os"
)

func Run() {
	input := internal.GetInput()

	udpAddr, err := net.ResolveUDPAddr("udp", os.Getenv("ADDRESS")+":"+os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open socket
	conection, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := conection.Write([]byte(input)); err != nil {
		fmt.Println("unable to write to connection")
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	_, _, err = conection.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(buffer))
}
