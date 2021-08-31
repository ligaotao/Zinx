package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	conn *net.TCPConn

	connID uint32

	isClosed bool

	handleApi ziface.HandleFunc

	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, callbackApi ziface.HandleFunc) *Connection {
	s := &Connection{
		conn:      conn,
		connID:    connId,
		isClosed:  false,
		handleApi: callbackApi,
		ExitChan:  make(chan bool),
	}
	return s
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")

	defer fmt.Println("链接ID", c.GetConnId(), "客户端地址", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.conn.Read(buf)

		if err != nil {
			fmt.Println("buf读取失败")
			continue
		}
		// 调用当前链接绑定的HandlerApi
		if err := c.handleApi(c.conn, buf, cnt); err != nil {
			fmt.Println("函数处理错误", err)
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn 开启 链接ID", c.GetConnId())
	// 启动当前链接的读数据的业务
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn 停止 链接ID", c.GetConnId())

	if c.isClosed {
		return
	}
	c.isClosed = true

	c.conn.Close()

	close(c.ExitChan)
}

func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnId() uint32 {
	return c.connID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) Send(body []byte) error {
	return nil
}
