package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	_ = rpc.RegisterName("HelloService", &HelloService{})
	http.HandleFunc("/rpcjson", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			w,
			r.Body,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))

	})
	http.ListenAndServe(":8080", nil)
}

func (s *HelloService) Hello(request string, reply *string) error {
	fmt.Println(request)
	*reply = "hello:" + request
	return nil

}

type HelloService struct {
}
