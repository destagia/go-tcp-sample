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
	clientMaxNum, err := strconv.ParseInt(os.Args[5], 10, 64)
	messagePerSec, err := strconv.ParseInt(os.Args[6], 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("FPS: %d\n", messagePerSec)
	fps := time.Duration(int(float64(1) / float64(messagePerSec) * float64(time.Second)))

	fmt.Printf("FPS sec: %d\n", fps)

	ch := make(chan int, 1000)

	for i := 0; i < int(clientNum); i++ {
		go client(serverIP, serverPort, int(deadline), fps, ch)
	}

	for i := 0; i < int(clientMaxNum - clientNum); i++ {
		if i % 100 == 0 {
			fmt.Printf("client: %d\n", int(clientNum) + i + 1)
			time.Sleep(time.Duration(3 * time.Second))

			deadCount := 0
			for {
				select {
				case _ = <-ch:
					deadCount++
					continue
				default:
					break
				}
				break
			}
	
			fmt.Printf("%d clients is dead\n", deadCount)
			i -= deadCount
		}

		time.Sleep(time.Duration(100 * time.Millisecond))
		go client(serverIP, serverPort, int(deadline), fps, ch)
	}

	fmt.Println("Wait for finishing...")

	for {
		time.Sleep(time.Duration(1 * time.Second))

		deadCount := 0
		for {
			select {
			case _ = <-ch:
				deadCount++
				continue
			default:
				break
			}
			break
		}

		// fmt.Printf("%d client is dead\n", deadCount)
	}
}

func client(serverIP string, serverPort string, deadline int, fps time.Duration, ch chan int) {

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverIP+":"+serverPort)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", tcpAddr.String())
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	for {
		// time.Sleep(time.Duration(fps))
		fmt.Printf(".")

		// conn.SetWriteDeadline(time.Now().Add(time.Duration(deadline) * time.Second))
		if _, err := conn.Write([]byte("hello world!")); err != nil {
			conn.Close()
			ch <- 0
			break
		}

		// conn.SetReadDeadline(time.Now().Add(time.Duration(deadline) * time.Second))
		if _, err := conn.Read(buf); err != nil {
			conn.Close()
			ch <- 0
			break
		}
	}
}