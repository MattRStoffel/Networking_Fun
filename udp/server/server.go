package udpserver

import (
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
)

func reverseWords(buf []byte) []byte {
	tmp := strings.Split(string(buf), " ")
	slices.Reverse(tmp)
	return []byte(strings.Join(tmp, " "))
}

func Run() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":4201")
	if err != nil {
		fmt.Println("unable to resolve address")
		os.Exit(1)
	}
	socket, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Cant hear you")
		os.Exit(1)
	}

	message := make([]byte, 1024)
	fmt.Println("Server running....")
	for {
		_, clientAddr, _ := socket.ReadFromUDP(message)
		words := reverseWords(message)
		fmt.Println(string(words))
		socket.WriteToUDP(words, clientAddr)
	}
}
