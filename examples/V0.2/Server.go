package main

import (
	IGerver "Gerver/Gerver/iface"
	Gerver "Gerver/Gerver/net"
	"fmt"
)

type PingRouter struct {
	Gerver.BaseRouter
}

func (br *PingRouter) Handle(req IGerver.IRequest) {
	data := req.GetData()
	fmt.Println(string(data))
	req.GetConnection().Send(data)
}

func main() {
	server := Gerver.NewServer()
	pingRouter := &PingRouter{}
	server.AddRouter(pingRouter)
	server.Start()
	fmt.Println("Server start")
	for {

	}
}
