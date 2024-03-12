package client_proxy

import (
	"learngo/new_helloword/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protcol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protcol, address) //建立连接
	if err != nil {
		panic("建立连接失败")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply) //调用远程方法
	if err != nil {
		return err
	}
	return nil
}
