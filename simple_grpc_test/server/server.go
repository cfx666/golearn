package main

import (
	"google.golang.org/grpc"
	"learngo/simple_grpc_test/handler"
	"learngo/simple_grpc_test/proto"
	"net"
)

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &handler.Server{})
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic("failed to listen")
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to serve")
	}
}
