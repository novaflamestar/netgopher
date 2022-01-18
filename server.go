package main

import (
	"fmt"
	"net"
)

type Server struct {
	Addr string
	Port string
}

func (s Server) Start() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", s.Addr+":"+s.Port)
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
	HandleConnection(conn)
}
