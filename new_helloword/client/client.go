package main

import (
	"fmt"
	"learngo/new_helloword/client_proxy"
)

func main() {

	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply *string = new(string)
	err := client.Hello("世界", reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(*reply)

}
