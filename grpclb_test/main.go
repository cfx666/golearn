package main

import (
	"context"
	"fmt"
	"learngo/grpclb_test/proto"
	"log"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"consul://127.0.0.1:8500/user_srv?wait=14s",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		userClient := proto.NewUserClient(conn)
		rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
			Pn:    1,
			PSize: 5,
		})
		if err != nil {
			panic(err)
		}

		for _, user := range rsp.Data {
			fmt.Println(user.NickName, user.Mobile, user.PassWord)
		}
	}
}
