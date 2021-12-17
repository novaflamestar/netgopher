package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	port, adder string
)

type argBundle struct {
	fType string
	adder string
	port  string
}

func parseArgs() (*argBundle, error) {
	c := flag.NewFlagSet("client", flag.ExitOnError)
	c.StringVar(&adder, "a", "", "Address to connect to")
	c.StringVar(&port, "p", "", "Port to connect to")

	s := flag.NewFlagSet("server", flag.ExitOnError)
	s.StringVar(&adder, "a", "0.0.0.0", "Address to listen on")
	s.StringVar(&port, "p", "", "Port to listen on")

	if len(os.Args) < 2 {
		return nil, fmt.Errorf("invalid usage, must specify client or server")
	}
	switch os.Args[1] {
	case "client":
		c.Parse(os.Args[2:])
		if port == "" || adder == "" {
			c.Usage()
			return nil, fmt.Errorf("invalid usage, client requires -a and -p flags")
		}
	case "server":
		s.Parse(os.Args[2:])
		if port == "" {
			s.Usage()
			return nil, fmt.Errorf("invalid usage, server requires the -p flag")
		}
	default:
		fmt.Println("Invalid usage, must specify client or server")
		c.Usage()
		s.Usage()
		os.Exit(1)
	}
	return &argBundle{os.Args[1], adder, port}, nil
}

func listen(a argBundle) {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", a.adder+":"+a.port)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Println("Server Successfully started on " + l.Addr().String())
	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received connection from %s\n", conn.RemoteAddr())
	handleConnection(conn)
}

func connect(a argBundle) {
	conn, err := net.Dial("tcp", a.adder+":"+a.port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to %s\n", conn.RemoteAddr())
	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
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

func main() {
	a, err := parseArgs()
	if err != nil {
		panic(err)
	}
	switch a.fType {
	case "server":
		listen(*a)
	case "client":
		connect(*a)
	default:
	}
}
