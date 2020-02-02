package main

import (
	"fmt"
	"net"
)

// Super simple first version with no error handling
// takes a tcp connection and return hello world


func main() {
	sock, _ := net.Listen("tcp", ":12345")

	
	for {
		conn, _ := sock.Accept()
		fmt.Fprintf(conn, "Hello World!\n\r")
		conn.Close()

	}
}
