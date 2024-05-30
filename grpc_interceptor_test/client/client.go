package main

import (
	"context"
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"learngo/grpc_interceptor_test/proto"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		//统计需要多长时间
		//开始时间
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		//结束时间
		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}

	var DialOpts []grpc.DialOption
	DialOpts = append(DialOpts, grpc.WithUnaryInterceptor(interceptor))
	DialOpts = append(DialOpts, grpc.WithInsecure())

	DialOpts = append(DialOpts, grpc.WithChainUnaryInterceptor(grpc_retry.UnaryClientInterceptor()))

	conn, err := grpc.Dial("localhost:8999", DialOpts...)
	if err != nil {
		panic("连接失败")
	}

	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	md := metadata.Pairs("name", "cui", "password", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "世界",
	}, grpc_retry.WithMax(3), grpc_retry.WithPerRetryTimeout(1*time.Second), grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded))
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
