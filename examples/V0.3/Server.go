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
	req.GetConnection().Send(data)
}

func main() {
	server := gnet.NewServer()
	pingRouter := &PingRouter{}
	server.AddRouter(pingRouter)
	server.Start()
	for {

	}
}
