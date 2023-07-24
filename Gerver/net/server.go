package Gerver

import (
	IGerver "Gerver/Gerver/iface"
	"fmt"
	"net"
)

type Server struct {
	router IGerver.IRouter
}

func NewServer() IGerver.IServer {
	s := &Server{
		router: nil,
	}
	return s
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	go func() {
		var ID uint64 = 0
		for {
			conn, err := listen.Accept() // 建立连接
			if err != nil {
				fmt.Println("accept failed, err:", err)
				continue
			}
			NC := NewConnection(s, conn, ID)
			NC.Start()
			ID++
		}
	}()
}

func (s *Server) AddRouter(router IGerver.IRouter) {
	s.router = router
}

func (s *Server) GetRouter() IGerver.IRouter {
	return s.router
}
