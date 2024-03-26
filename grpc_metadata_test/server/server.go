package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learngo/grpc_metadata_test/proto"
	"net"
)

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &HelloServer{})
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		panic("failed to listen")
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve")
	}
}

type HelloServer struct {
}

func (s *HelloServer) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}
	for key, value := range md {
		fmt.Println(key, value)
	}
	//打印的有协议头的相关信息，我们只需要取name和password，这就是map的用法
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}
