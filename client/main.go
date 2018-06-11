package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
)

func main() {
	fmt.Println("[benchmark client] start")

	serverIP := os.Args[1]
	serverPort := os.Args[2]
	deadline, err := strconv.ParseInt(os.Args[3], 10, 64)
	clientNum, err := strconv.ParseInt(os.Args[4], 10, 64)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(clientNum); i++ {
		go client(serverIP, serverPort, int(deadline))
	}

	var s string
	fmt.Scan(&s)
}

func client(serverIP string, serverPort string, deadline int) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", tcpAddr.String())
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Duration(33 * time.Millisecond))

		conn.SetWriteDeadline(time.Now().Add(time.Duration(deadline) * time.Second))
		conn.Write([]byte("hello world!"))

		buf := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(time.Duration(deadline) * time.Second))
		_, err := conn.Read(buf)

		if err != nil {
			panic(err)
		}
	}
}