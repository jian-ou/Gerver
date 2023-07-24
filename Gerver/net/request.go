package Gerver

import IGerver "Gerver/Gerver/iface"

type Request struct {
	dat        []byte
	connection IGerver.IConnection
}

func NewRequest(connection IGerver.IConnection, dat []byte) IGerver.IRequest {
	r := &Request{
		dat:        dat,
		connection: connection,
	}
	return r
}

func (r *Request) GetData() []byte {
	return r.dat
}

func (r *Request) GetConnection() IGerver.IConnection {
	return r.connection
}
