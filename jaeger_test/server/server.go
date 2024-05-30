package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/jaeger_test/proto"
	"net"
)

func main() {

	server := grpc.NewServer()
	proto.RegisterGreeterServer(server, &helloServer{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}

type helloServer struct{}

func (h *helloServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello，" + req.Name,
	}, nil
}
