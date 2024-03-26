package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/simple_grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "世界",
	})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
