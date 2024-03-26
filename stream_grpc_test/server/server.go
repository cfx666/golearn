package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learngo/stream_grpc_test/proto"
	"net"
	"sync"
	"time"
)

const PROT = ":60479"

type Server struct {
}

func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		_ = res.Send(&proto.StreamResData{Data: fmt.Sprintf("第%d次发送，time是%v", i, time.Now().Unix())})
		time.Sleep(time.Second)
		i++
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *Server) PutStream(clistr proto.Greeter_PutStreamServer) error {
	for {
		req, err := clistr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(req.Data)
	}
	return nil
}

func (s *Server) AllStream(allstr proto.Greeter_AllStreamServer) error {

	var wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			_ = allstr.Send(&proto.StreamResData{Data: fmt.Sprintf("我是服务器%d", i)})
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			req, err := allstr.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("收到客户端消息：" + req.Data)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PROT)
	if err != nil {
		panic("failed to listen")
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		panic("failed to serve")
	}
}
