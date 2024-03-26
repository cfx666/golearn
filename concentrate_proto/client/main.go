package main

import (
	"fmt"
	"learngo/concentrate_proto/proto"
)

func main() {
	var req = proto.HelloRequest{
		Name: "cui",
		Data: []string{"a", "b", "c"},
	}
	fmt.Println(req)

}
