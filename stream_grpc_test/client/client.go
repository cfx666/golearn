package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learngo/stream_grpc_test/proto"
)

func main() {

	conn, err := grpc.Dial("localhost:60479", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// 服务端流模式
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "世界"})
	for {
		r, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		println(r.Data)
	}

	// 客户端流模式
	/*putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		_ = putS.Send(&grpc_validate_test.proto.StreamReqData{Data: fmt.Sprintf("第%d次发送", i)})
		time.Sleep(time.Second)
		i++
		if i > 10 {
			break
		}
	}*/

	// 双向流模式
	/*allstr, _ := c.AllStream(context.Background())
	var wg = sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 0
		for {
			i++
			_ = allstr.Send(&proto.StreamReqData{Data: fmt.Sprintf("我是客户端%d", i)})
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
			fmt.Println("收到服务器的消息：" + req.Data)
		}
	}()
	wg.Wait()*/

}
