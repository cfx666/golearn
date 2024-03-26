package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/grpc_interceptor_test/proto"
	"net"
)

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		//使用1
		//拦截到请求之后，可以在这里进行一些处理
		println("before")
		//处理完成之后，调用原来的处理逻辑
		return handler(ctx, req)

		//使用2：想运行原来的处理逻辑以后，再进行一些处理
		println("before")
		resp, err = handler(ctx, req)
		println("after")
		return
	}
	serOpt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(serOpt)
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
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}
