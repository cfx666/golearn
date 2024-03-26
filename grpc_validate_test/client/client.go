package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/grpc_validate_test/proto"
)

type customCredential struct{}

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	rsp, err := c.SayHello(context.Background(), &proto.Person{
		Id:     200000,
		Email:  "cui@cui.com",
		Mobile: "18934566666",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id)
}
