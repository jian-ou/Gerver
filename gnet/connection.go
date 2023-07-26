package gnet

import (
	"Gerver/gcoder"
	"Gerver/gconf"
	"Gerver/giface"
	"bufio"
	"fmt"
	"net"
)

type Connection struct {
	server giface.IServer
	conn   net.Conn
	connID uint64

	state uint8

	buf     []byte
	coder   giface.ICoder
	bufSize int

	PreHandle  func(giface.IConnection)
	PostHandle func(giface.IConnection)

	msgBufChan chan []byte
	closeChan  chan bool
}

func NewConnection(server giface.IServer, conn net.Conn, connID uint64, state uint8) giface.IConnection {
	c := &Connection{
		conn:       conn,
		connID:     connID,
		msgBufChan: make(chan []byte),
		server:     server,
		state:      0,
		bufSize:    gconf.Globalconf.MaxConnBufSize,
		buf:        make([]byte, 0),
		closeChan:  make(chan bool, 1),
		coder:      gcoder.NewTLVCoder(),
	}
	c.SetState(state)
	c.PreHandle = server.GetPreHandle()
	c.PostHandle = server.GetPostHandle()
	return c
}

func (c *Connection) StartReader() {
	for {
		reader := bufio.NewReader(c.conn)
		buf := make([]byte, c.bufSize)
		n, err := reader.Read(buf[:])
		if err != nil {
			c.closeChan <- true
			// fmt.Println("read from client failed, err:", err)
			break
		}
		c.buf = append(c.buf, buf[:n]...)
		for {
			msgid, length, dat, err := c.coder.Decode(c.buf[:])
			if err != nil {
				break
			}
			NR := NewRequest(c, dat, (c.GetServer().GetRouter(msgid)))
			c.GetServer().GetDispatch().AddRequest(NR)

			c.buf = c.buf[length:]
			if len(c.buf) == 0 {
				break
			}
		}
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
	if err := c.conn.Close(); err != nil {
		fmt.Println("Conn Close err : ", err)
	}
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

func (c *Connection) SetState(state uint8) {
	c.state = state
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}
