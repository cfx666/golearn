package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	dial, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	var reply string
	err = dial.Call("HelloService.Hello", "世界", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

}
