package gnet

import "Gerver/giface"

type Request struct {
	dat        []byte
	connection giface.IConnection
}

func NewRequest(connection giface.IConnection, dat []byte) giface.IRequest {
	r := &Request{
		dat:        dat,
		connection: connection,
	}
	return r
}

func (r *Request) GetData() []byte {
	return r.dat
}

func (r *Request) GetConnection() giface.IConnection {
	return r.connection
}
