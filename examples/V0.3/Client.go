package main

import (
	"Gerver/gnet"
	"time"
)

func main() {
	client := gnet.NewClient()
	client.Start()
	for {
		client.Send([]byte("hello"))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
