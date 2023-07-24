package main

import (
	Glient "Gerver/Glient/net"
	"time"
)

func main() {
	client := Glient.NewClient()
	client.Start()
	for {
		client.Send([]byte("hello"))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
