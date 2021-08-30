package znet

import (
	"Zinx/ziface"
	"fmt"
	"net"
)

type Server struct {
	// 名称
	Name string
	// ip
	IPVersion string
	IP string
	// 端口
	Port int
}

func (s *Server) Start () {

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

		for {
			conn,err := listener.AcceptTCP()

			if err != nil {
				fmt.Println("Accept Err", err)
				continue
			}
			// 已经与客户端建立了链接

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("读取buf错误", err)
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("回写错误")
						continue
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop () {

}

func (s *Server) Serve () {
	s.Start()
	select {}
}

/*
*	初始化Server方法
*/

func NewServer(name string) ziface.IServer  {
	s := &Server{
		Name: name,
		IPVersion: "tcp4",
		Port: 8000,
		IP: "0.0.0.0",
	}
	return s
}