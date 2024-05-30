package main

import (
	"net"
	"net/rpc"
)

func main() {

	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	rpc.Accept(listen)

}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
