package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	newProducer, err := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.137.129:9876"}))
	if err != nil {
		panic("create producer error")
	}

	err = newProducer.Start()
	if err != nil {
		panic("start producer error")
	}

	res, err := newProducer.SendSync(context.Background(), &primitive.Message{
		Topic: "test",
		Body:  []byte("Hello RocketMQ Go Client!"),
	})
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}

	if err = newProducer.Shutdown(); err != nil {
		fmt.Printf("shutdown producer error: %s", err)
	}
}
