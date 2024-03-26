package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/grpc_auth_interceptor_test/proto"
)

type customCredential struct {
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "10133",
		"appkey": "I am key",
	}, nil
}
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {

	DialOpt := grpc.WithPerRPCCredentials(customCredential{})
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure(), DialOpt)

	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "世界",
	})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.Message)
}
