package znet

import (
	"Zinx/ziface"
	"errors"
	"fmt"
	"net"
)

type Server struct {
	// 名称
	Name string
	// ip
	IPVersion string
	IP        string
	// 端口
	Port int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("链接绑定函数被调用")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("回写错误")
		return errors.New("callback error")
	}
	return nil
}

func (s *Server) Start() {

	go func() {
		//1. 获取一个TCP Addr
		fmt.Printf("服务器启动 监听IP: %v 端口 %d", s.IP, s.Port)
		//2 监听服务器地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("地址解析错误")
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("监听失败", err)
		}
		fmt.Println("Zinx启动成功")
		//3. 阻塞等待客户端链接
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()

			if err != nil {
				fmt.Println("Accept Err", err)
				continue
			}
			// 已经与客户端建立了链接
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	select {}
}

/*
*	初始化Server方法
 */

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		Port:      8000,
		IP:        "0.0.0.0",
	}
	return s
}
