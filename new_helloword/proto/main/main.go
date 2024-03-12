package main

import (
	"encoding/json"
	"fmt"
	p "github.com/golang/protobuf/proto"
	"learngo/new_helloword/proto"
)

type Hello struct {
	Name string `json:"name"`
}

func main() {
	req := proto.HelloRequest{
		Name: "cui",
	}
	rsp, _ := p.Marshal(&req) //序列化，返回[]byte。具体是如何做到的，可以搜索一下protobuf的原理 varint
	fmt.Println(rsp)          //[10 3 99 117 105]
	fmt.Println(string(rsp))
	fmt.Println(len(rsp)) //5

	//和json对比一下
	newreq, _ := json.Marshal(Hello{
		Name: "cui",
	})
	fmt.Println(string(newreq))
	fmt.Println(len(newreq)) //14

	//反序列化
	req2 := proto.HelloRequest{}
	_ = p.Unmarshal(rsp, &req2)
	fmt.Println(req2.Name) //cui
}
