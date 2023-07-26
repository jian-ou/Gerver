package giface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetConn() net.Conn
	GetConnID() uint64
	GetServer() IServer
	Send([]byte)
	SetState(uint8)
	RemoteAddr() net.Addr //获取链接远程地址信息
	LocalAddr() net.Addr  //获取链接本地地址信息
}
