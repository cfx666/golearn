package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learngo/grpc_metadata_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8999", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	md := metadata.Pairs("name", "cui", "password", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "世界",
	})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
