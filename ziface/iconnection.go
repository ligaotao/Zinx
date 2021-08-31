package ziface

import "net"

type IConnection interface {
	// Start 启动链接
	Start()
	// Stop 结束链接
	Stop()
	// GetTcpConnection 获取当前连接的socket conn
	GetTcpConnection() *net.TCPConn
	// GetConnId 获取链接的Id
	GetConnId() uint32
	// RemoteAddr 获取客户端的TCP状态
	RemoteAddr() net.Addr
	// Send 发送数据
	Send(data []byte) error
}

// 定义一个处理链接业务的方法

type HandleFunc func(*net.TCPConn, []byte, int) error
