package main

import (
	"Gerver/giface"
	"Gerver/gnet"
	"fmt"
)

type PingRouter struct {
	gnet.BaseRouter
}

func (br *PingRouter) Handle(req giface.IRequest) {
	data := req.GetData()
	fmt.Println(string(data))
	req.GetConnection().Send(append(data, []byte(" Ping....")...))
}

type GoodRouter struct {
	gnet.BaseRouter
}

func (br *GoodRouter) Handle(req giface.IRequest) {
	data := req.GetData()
	fmt.Println(string(data))
	req.GetConnection().Send(append(data, []byte(" Good....")...))
}

func main() {
	server := gnet.NewServer()
	server.AddRouter(100, &PingRouter{})
	server.AddRouter(200, &GoodRouter{})
	server.Start()
	for {

	}
}
