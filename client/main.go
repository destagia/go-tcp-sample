package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("[benchmark client] start")

	serverIP := os.Args[1]
	serverPort := os.Args[2]

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", tcpAddr.String())
	if err != nil {
		panic(err)
	}

	count := 0

	for {
		time.Sleep(time.Duration(1000 * time.Millisecond))

		conn.Write([]byte("hello world!"))

		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		count++
		fmt.Printf("\r%d", count)
	}
}