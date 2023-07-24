package gnet

import (
	"Gerver/gconf"
	"Gerver/giface"
	"fmt"
	"net"
)

type Server struct {
	Name     string
	Version  string
	HostPort uint
	router   giface.IRouter
}

func NewServer() giface.IServer {
	s := &Server{
		router: nil,
	}
	s.Name = gconf.Globalconf.Name
	s.Version = gconf.Globalconf.Version
	s.HostPort = gconf.Globalconf.HostPort
	return s
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", (fmt.Sprintf("127.0.0.1:%d", s.HostPort)))
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Printf("[%s Server] start\n", s.Name)
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

func (s *Server) AddRouter(router giface.IRouter) {
	s.router = router
}

func (s *Server) GetRouter() giface.IRouter {
	return s.router
}
