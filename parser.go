package main

import (
	"flag"
	"fmt"
	"os"
)

type argBundle struct {
	FType string
	Adder string
	Port  string
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
