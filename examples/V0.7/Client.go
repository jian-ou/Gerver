package main

import (
	"Gerver/gcoder"
	"Gerver/gnet"
	"time"
)

func main() {
	client := gnet.NewClient()
	client.Start()
	c := gcoder.NewTLVCoder()
	for {
		client.Send(c.Encode(100, []byte("hello world")))
		client.Send(c.Encode(200, []byte("hello world")))
		time.Sleep(time.Duration(1) * time.Microsecond)
		client.Send(c.Encode(200, []byte("hello world")))
		time.Sleep(time.Duration(1) * time.Microsecond)
		client.Send(c.Encode(201, []byte("hello world")))
		time.Sleep(time.Duration(1) * time.Microsecond)
	}
}
