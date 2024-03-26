package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"learngo/grpc_auth_interceptor_test/proto"
	"net"
)

func main() {
	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无Token认证信息")
		}
		fmt.Printf("metadata: %v\n", md)
		var (
			appid  string
			appkey string
		)
		if value, ok := md["appid"]; ok {
			appid = value[0]
		}
		if value, ok := md["appkey"]; ok {
			appkey = value[0]
		}
		if appid != "101" || appkey != "I am key" {
			return resp, status.Error(codes.Unauthenticated, "Token认证信息无效")
		}

		//处理完成之后，调用原来的处理逻辑
		return handler(ctx, req)
	}
	serOpt := grpc.UnaryInterceptor(interceptor)
	g := grpc.NewServer(serOpt)
	proto.RegisterGreeterServer(g, &HelloServer{})
	lis, err := net.Listen("tcp", ":9001")
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
