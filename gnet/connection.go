package gnet

import (
	"Gerver/gcoder"
	"Gerver/giface"
	"bufio"
	"fmt"
	"net"
)

type Connection struct {
	server giface.IServer
	conn   net.Conn
	connID uint64

	msgBufChan chan []byte
}

func NewConnection(server giface.IServer, conn net.Conn, connID uint64) giface.IConnection {
	c := &Connection{
		conn:       conn,
		connID:     connID,
		msgBufChan: make(chan []byte),
		server:     server,
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader start", c.GetConnID())
	for {
		reader := bufio.NewReader(c.conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		cr := gcoder.NewTLVDecoder()
		msgid, _, dat := cr.Decode(buf[:n])
		NR := NewRequest(c, dat, (c.GetServer().GetRouter(msgid)))
		NR.Run()
		// c.GetServer().GetRouter().PreHandle(NR)
		// c.GetServer().GetRouter().Handle(NR)
		// c.GetServer().GetRouter().PostHandle(NR)
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("Writer start", c.GetConnID())
	for {
		select {
		case data, ok := <-c.msgBufChan:
			if ok {
				fmt.Println("get data : ", data)
				c.Send(data)
			}
		}
	}
}

func (c *Connection) Start() {
	go c.StartReader()
	go c.StartWriter()
}
func (c *Connection) Stop() {

}

func (c *Connection) GetConnID() uint64 {
	return c.connID
}

func (c *Connection) Send(data []byte) {
	_, err := c.conn.Write(data)
	if err != nil {
		fmt.Println("Send err :", err)
	}
}

func (c *Connection) GetServer() giface.IServer {
	return c.server
}

func (c *Connection) GetConn() net.Conn {
	return c.conn
}
