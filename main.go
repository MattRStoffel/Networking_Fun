package main

import (
	"NetworkingFun/tcp/client"
	"NetworkingFun/tcp/server"
	"NetworkingFun/udp/client"
	"NetworkingFun/udp/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

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
			if err := tcpclient.Run(addr, port); err != nil {
				fmt.Println(err)
			}
		case "server", "s", "S", "Server":
			log.Println("Starting server...")
			if err := tcpserver.Run(port); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Invalid seletion")
		}
	case "U", "u", "udp", "Udp", "":
		switch mode {
		case "Client", "C", "client", "c":
			log.Println("Starting client...")
			if err := udpclient.Run(addr, port); err != nil {
				fmt.Println(err)
			}
		case "server", "s", "S", "Server", "":
			log.Println("Starting server...")
			if err := udpserver.Run(port); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Invalid seletion")
		}
	default:
		fmt.Println("Invalid Protocol Selection")
	}
}
