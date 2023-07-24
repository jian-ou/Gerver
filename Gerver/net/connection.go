package Gerver

import (
	IGerver "Gerver/Gerver/iface"
	"bufio"
	"fmt"
	"net"
)

type Connection struct {
	conn       net.Conn
	connID     uint64
	server     IGerver.IServer
	msgBufChan chan []byte
}

func NewConnection(server IGerver.IServer, conn net.Conn, connID uint64) IGerver.IConnection {
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
		NR := NewRequest(c, buf[:n])
		c.GetServer().GetRouter().PreHandle(NR)
		c.GetServer().GetRouter().Handle(NR)
		c.GetServer().GetRouter().PostHandle(NR)
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

func (c *Connection) GetServer() IGerver.IServer {
	return c.server
}

func (c *Connection) GetConn() net.Conn {
	return c.conn
}
