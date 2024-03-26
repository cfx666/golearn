package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"learngo/grpc_error_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	//grpc的超时机制
	/*ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	_, err = c.SayHello(ctx, &proto.HelloRequest{
		Name: "世界",
	})*/

	//grpc的错误处理，之前处理错误是直接panic
	_, err = c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "世界",
	})
	/*if err != nil {
		panic("调用失败")
	}*/
	s, ok := status.FromError(err)
	if ok {
		fmt.Println(s.Message())
		fmt.Println(s.Code())
	}
}
