package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type NetConnectors interface {
	Start()
}

func HandleConnection(conn net.Conn) {
	in, out := make(chan string), make(chan string)
	quit := make(chan bool)
	go startSender(conn, out, quit)
	go startReceiver(conn, in, quit)
	for {
		select {
		case i := <-in:
			fmt.Print(i)
		case o := <-out:
			conn.Write([]byte(o))
		case <-quit:
			return
		}
	}
}

func startSender(conn net.Conn, o chan string, q chan bool) {
	r := bufio.NewReader(os.Stdin)
	for {
		message, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("Closing connection to %s\n", conn.RemoteAddr())
			conn.Close()
			q <- true
			return
		}
		o <- message
	}
}

func startReceiver(conn net.Conn, i chan string, q chan bool) {
	for {
		message, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			fmt.Printf("Closing connection to %s\n", conn.RemoteAddr())
			conn.Close()
			q <- true
			return
		}
		i <- string(message)
	}
}
