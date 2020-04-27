/*
 * A baisc client that will have similar functionality as curl
 * 
 */




package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main () {

	// TO DO Learn how to  take comand line input
	var addr string = os.Args[1]

	conn, _ := net.Dial("tcp", addr)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(message)


}
