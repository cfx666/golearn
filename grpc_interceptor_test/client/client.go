package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"learngo/grpc_interceptor_test/proto"
	"time"
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
	DialOpt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial("localhost:8999", grpc.WithInsecure(), DialOpt)
	//改变
	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts, grpc.WithInsecure())
	dialOpts = append(dialOpts, DialOpt)
	//conn, err := grpc.Dial("localhost:8999", dialOpts...)

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
