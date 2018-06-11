package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := os.Args[1]
	fmt.Printf("[benchmark server] start [:%v]\n", port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%s", port))
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
		fmt.Printf("accepted: %s", conn.RemoteAddr().String())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	buf := make([]byte, 256)
	for {
		var mes []byte
		if n, err := conn.Read(buf); err != nil {
			fmt.Println(err.Error())
			break
		} else {
			mes = buf[:n]
		}

		if _, err := conn.Write(mes); err != nil {
			fmt.Println(err.Error())
		}
	}

	conn.Close()
}