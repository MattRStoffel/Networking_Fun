package main

import (
	"NetworkingFun/tcp/client"
	"NetworkingFun/tcp/server"
	"NetworkingFun/udp/client"
	"NetworkingFun/udp/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var protocol string
	fmt.Println("select protocol ([U]dp, [t]cp): ")
	fmt.Scanln(&protocol)

	var mode string
	fmt.Print("select mode ([S]erver, [c]lient): ")
	fmt.Scanln(&mode)

	switch protocol {
	case "t", "T", "tcp", "Tcp":
		switch mode {
		case "Client", "C", "client", "c", "":
			log.Println("Starting client...")
			tcpclient.Run()
		case "server", "s", "S", "Server":
			log.Println("Starting server...")
			tcpserver.Run()
		default:
			fmt.Println("Invalid seletion")
		}
	case "U", "u", "udp", "Udp", "":
		switch mode {
		case "Client", "C", "client", "c":
			log.Println("Starting client...")
			udpclient.Run()
		case "server", "s", "S", "Server", "":
			log.Println("Starting server...")
			udpserver.Run()
		default:
			fmt.Println("Invalid seletion")
		}
	default:
		fmt.Println("Invalid Protocol Selection")
	}
}
