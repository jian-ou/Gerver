package IGerver

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConn() net.Conn
	GetConnID() uint64
	GetServer() IServer
	Send([]byte)
}
