package main

import "Zinx/znet"

func main()  {
	s := znet.NewServer("V0.01")
	s.Serve()
}
