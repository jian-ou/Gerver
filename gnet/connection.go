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

	state uint16

	PreHandle  func(giface.IConnection)
	PostHandle func(giface.IConnection)

	msgBufChan chan []byte
	closeChan  chan bool
}

func NewConnection(server giface.IServer, conn net.Conn, connID uint64, state uint16) giface.IConnection {
	c := &Connection{
		conn:       conn,
		connID:     connID,
		msgBufChan: make(chan []byte),
		server:     server,
		state:      0,
		closeChan:  make(chan bool, 1),
	}
	c.SetState(state)
	c.PreHandle = server.GetPreHandle()
	c.PostHandle = server.GetPostHandle()
	return c
}

func (c *Connection) StartReader() {
	for {
		reader := bufio.NewReader(c.conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			c.closeChan <- true
			// fmt.Println("read from client failed, err:", err)
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
	for {
		select {
		case data, ok := <-c.msgBufChan:
			if ok {
				fmt.Println("get data : ", data)
				c.Send(data)
			}
		case b := <-c.closeChan:
			if b {
				c.Stop()
				return
			}
		}
	}
}

func (c *Connection) Start() {
	c.PreHandle(c)

	go c.StartReader()
	go c.StartWriter()
}
func (c *Connection) Stop() {
	c.PostHandle(c)
	close(c.closeChan)
	close(c.msgBufChan)
	c.SetState(0)
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

func (c *Connection) SetState(state uint16) {
	c.state = state
}
