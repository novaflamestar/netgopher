package main

import (
	"fmt"
	"net"
)

type Client struct {
	Addr string
	Port string
}

func (c Client) Start() {
	conn, err := net.Dial("tcp", c.Addr+":"+c.Port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to %s\n", conn.RemoteAddr())
	HandleConnection(conn)
}
