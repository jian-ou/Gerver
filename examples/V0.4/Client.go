package main

import (
	"Gerver/gcoder"
	"Gerver/gnet"
	"time"
)

func main() {
	client := gnet.NewClient()
	client.Start()
	c := gcoder.NewTLVDecoder()
	for {
		client.Send(c.Encode(100, []byte("hello world")))
		time.Sleep(time.Duration(1) * time.Second)
		client.Send(c.Encode(200, []byte("hello world")))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
