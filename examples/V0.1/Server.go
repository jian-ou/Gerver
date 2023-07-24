package main

import (
	Gerver "Gerver/Gerver/net"
	"fmt"
)

func main() {
	server := Gerver.NewServer()
	server.Start()
	fmt.Println("Server start")
	for {

	}
}
