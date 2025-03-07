package udpclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Run() {
	// Get input from user
	fmt.Println("input a lowwercase sentence: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	input = strings.Split(input, "\n")[0]

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

	buffer := make([]byte, 1024)
	_, _, err = conection.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(buffer))

}
