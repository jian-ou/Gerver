package gnet

import (
	"Gerver/gconf"
	"Gerver/giface"
	"fmt"
	"net"
)

type Client struct {
	conn     net.Conn
	Name     string
	Version  string
	Host     string
	HostPort uint
}

func NewClient() giface.IClient {
	c := &Client{
		conn: nil,
	}
	c.Name = gconf.Globalconf.Name
	c.Version = gconf.Globalconf.Version
	c.Host = gconf.Globalconf.Host
	c.HostPort = gconf.Globalconf.HostPort
	return c
}

func (c *Client) Start() {
	var err error
	c.conn, err = net.Dial("tcp", (fmt.Sprintf("127.0.0.1:%d", c.HostPort)))
	if err != nil {
		fmt.Println("Dial err :", err)
		return
	}
	fmt.Printf("[%s Client] start\n", c.Name)
	buf := make([]byte, 128)
	go func() {
		for {
			n, err := c.conn.Read(buf[:])
			if err != nil {
				fmt.Println("recv failed, err:", err)
				return
			}
			fmt.Println(string(buf[:n]))
		}
	}()
}

func (c *Client) Send(b []byte) error {
	_, err := c.conn.Write(b) // 发送数据
	if err != nil {
		return err
	}
	return nil
}
