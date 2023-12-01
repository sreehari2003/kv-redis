package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Database struct {
	data map[string]string
}

// After running the code you can access the code via telnet
// eg-> telnet localhost 8080
// 8080 is the port u give when u run the code

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
	// server will close on cleanup
	defer server.Close()

	// infinite loop
	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		db := Database{
			data: make(map[string]string),
		}

		/*
			         When connections are establised from each port the code via TELNET
					 concurrently new threads are created for each user
		*/
		go handleConnection(client, &db)
	}
}

func handleConnection(s net.Conn, db *Database) {
	fmt.Printf("Serving %s", s.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(s).ReadString('\n')
		if err != nil {
			fmt.Println("error reading:", err)
			break
		}
		temp := strings.Fields(netData)

		if len(temp) == 0 {
			continue
		}

		command := temp[0]
		key := temp[1]
		value := temp[2]

		/*
		   This code block is used to show some  response to the end user
		   s.Write([]byte("OK\n"))
		*/
		switch command {
		case "SET":
			db.set(key, value)
		case "HELP":
			s.Write([]byte(`
SET key value: used to set a key , value to Database,
GET key : used to get the key, value from already existing db
DELETE key: Used to delete the key , value from database
`))

		default:
			s.Write([]byte("Invalid Input\n"))
		}

	}
	fmt.Println("Closing client")
	s.Close()
}

func (s *Database) set(key, value string) {
	s.data[key] = value
}
