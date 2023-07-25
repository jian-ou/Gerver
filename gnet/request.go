package gnet

import "Gerver/giface"

type Request struct {
	connection giface.IConnection
	dat        []byte
	router     giface.IRouter
}

func NewRequest(connection giface.IConnection, dat []byte, router giface.IRouter) giface.IRequest {
	r := &Request{
		dat:        dat,
		connection: connection,
		router:     router,
	}
	return r
}

func (r *Request) GetData() []byte {
	return r.dat
}

func (r *Request) GetConnection() giface.IConnection {
	return r.connection
}

func (r *Request) Run() {
	r.router.PreHandle(r)
	r.router.Handle(r)
	r.router.PostHandle(r)
}
