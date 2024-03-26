package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learngo/grpc_validate_test/proto"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.Person) (*proto.Person,
	error) {
	return &proto.Person{
		Id: 32,
	}, nil
}

type Validator interface { //为什么要定义一个接口，因为我们要在拦截器中判断哪些结构体实现了这个接口，因为拦截器是拦截全部的请求，有些参数没有实现校验接口。如果实现了，就调用Validate方法，这样就可以实现对请求参数的校验
	Validate() error
}

func main() {

	interceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if r, ok := req.(Validator); ok { //判断请求参数是否实现了校验接口，这里也可以使用proto.Person，但是这样就只能校验Person结构体了，如果有其他结构体也需要校验，就需要再写一遍，所以这里使用接口
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	g := grpc.NewServer(opts...)
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
