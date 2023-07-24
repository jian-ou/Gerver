package main

import (
	Glient "Gerver/Glient/net"
	"fmt"
	"time"
)

func main() {
	client := Glient.NewClient()
	client.Start()
	fmt.Println("Client start")
	for {
		client.Send([]byte("hello"))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
