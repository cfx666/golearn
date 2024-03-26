package main

import (
	"learngo/new_helloword/handler"
	"learngo/new_helloword/server_proxy"
	"net"
	"net/rpc"
)

func main() {

	listener, _ := net.Listen("tcp", ":1234")
	_ = server_proxy.RegisterHelloService(&handler.NewHelloService{})
	for {
		conn, _ := listener.Accept() //接收请求
		go rpc.ServeConn(conn)       //处理请求，所有的请求都由rpc.ServeConn处理
	}
}
