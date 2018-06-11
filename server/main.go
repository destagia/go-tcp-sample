package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := os.Args[1]
	fmt.Printf("[benchmark server] start [:%v]\n", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		conn.Write(buf[:n])
	}
}