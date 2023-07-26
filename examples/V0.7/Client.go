package main

import (
	"Gerver/gcoder"
	"Gerver/giface"
	"Gerver/gnet"
	"time"
)

var c []giface.IClient

func main() {
	c := make([]giface.IClient, 100)
	for i := 0; i < 100; i++ {
		c[i] = gnet.NewClient()
		c[i].Start()
	}
	coder := gcoder.NewTLVCoder()
	time.Sleep(time.Duration(5) * time.Second)
	for {
		for i := 0; i < 100; i++ {
			c[i].Send(coder.Encode(100, []byte("hello world")))
		}
		time.Sleep(time.Duration(1) * time.Microsecond)
	}
}
