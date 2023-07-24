package Glient

import (
	Glient "Gerver/Glient/iface"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient() Glient.IClient {
	c := &Client{
		conn: nil,
	}
	return c
}

func (c *Client) Start() {
	var err error
	c.conn, err = net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial err :", err)
		return
	}
	go func() {
		for {
			buf := [512]byte{}
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
