package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port")
		return
	}
	port := ":" + arguments[1]

	server, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connected to tcp server on port " + arguments[1])
	defer server.Close()

	// infinite loop
	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// go co-routines
		go handleConnection(client)
	}
}

func handleConnection(s net.Conn) {
	defer s.Close()
	fmt.Printf("Serving %s", s.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(s).ReadString('\n')
		if err != nil {
			fmt.Println("error reading:", err)
			break
		}
		temp := strings.TrimSpace(netData)
		fmt.Println("recived: ", temp)
		s.Write([]byte(temp + "\n"))
	}
	fmt.Println("Closing client")
	s.Close()
}
