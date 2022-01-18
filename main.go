package main

import (
	"fmt"
	"os"
)

var (
	port, adder string
)

func main() {
	app := App{}
	a, err := parseArgs()
	app.SetArgs(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	app.SetUpNetConn()
	app.Net.Start()
}
