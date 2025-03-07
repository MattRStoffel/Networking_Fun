package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Get input from user
	fmt.Println("input a lowwercase sentence: ")
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		fmt.Println("Bad input")
		os.Exit(1)
	}

	address := "localhost:4201"
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("unable to resolve address")
		os.Exit(1)
	}

	// Open socket
	conection, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Socket open fail")
		os.Exit(1)
	}

	if _, err := conection.Write([]byte(input)); err != nil {
		fmt.Println("unable to write to connection")
		os.Exit(1)
	}

	message := make([]byte, 1024)
	n, _, err := conection.ReadFromUDP(message)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(message[:n]))

}
