package gnet

import (
	"Gerver/gconf"
	"Gerver/giface"
	"Gerver/glogo"
	"fmt"
	"net"
)

type Server struct {
	Name     string
	Version  string
	HostPort uint

	PreHandle  func(giface.IConnection)
	PostHandle func(giface.IConnection)

	// connections map[uint32]giface.IConnection

	// manage    giface.IManage
	dispatch giface.IDispatch
	routers  map[uint32]giface.IRouter
}

func NewServer() giface.IServer {
	s := &Server{
		routers: make(map[uint32]giface.IRouter),
	}
	s.Name = gconf.Globalconf.Name
	s.Version = gconf.Globalconf.Version
	s.HostPort = gconf.Globalconf.HostPort
	s.PreHandle = func(i giface.IConnection) {}
	s.PostHandle = func(i giface.IConnection) {}
	s.dispatch = NewDispatch(s, gconf.Globalconf.MaxProcess)
	return s
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", (fmt.Sprintf("127.0.0.1:%d", s.HostPort)))
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	glogo.Logo()
	fmt.Printf("[%s-Server] start at [%s:%d]\n", s.Name, "0.0.0.0", s.HostPort)
	go func() {
		var ID uint64 = 0
		for {
			conn, err := listen.Accept() // 建立连接
			if err != nil {
				fmt.Println("accept failed, err:", err)
				continue
			}
			NC := NewConnection(s, conn, ID, 1)
			NC.Start()
			ID++
		}
	}()
}

func (s *Server) AddRouter(msgID uint32, router giface.IRouter) {
	s.routers[msgID] = router
}

func (s *Server) GetRouter(msgID uint32) giface.IRouter {
	if s.routers[msgID] != nil {
		return s.routers[msgID]
	}
	return &NoneRouter{}
}

func (s *Server) AddPreHandle(p func(giface.IConnection)) {
	s.PreHandle = p
}
func (s *Server) AddPostHandle(p func(giface.IConnection)) {
	s.PostHandle = p
}

func (s *Server) GetPreHandle() func(giface.IConnection) {
	return s.PreHandle
}
func (s *Server) GetPostHandle() func(giface.IConnection) {
	return s.PostHandle
}

func (s *Server) GetDispatch() giface.IDispatch {
	return s.dispatch
}
