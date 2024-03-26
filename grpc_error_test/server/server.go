package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/grpc_error_test/proto"
	"net"
	"time"
)

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("failed to listen")
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve")
	}
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {

	//错误处理
	//return nil, status.Errorf(codes.InvalidArgument, "invalid argument %s", request.Name)

	//超时机制
	time.Sleep(5 * time.Second)
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}
