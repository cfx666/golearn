package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"time"
)

func main() {
	pushConsumer, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.137.129:9876"}),
		consumer.WithGroupName("lpx"))

	if err := pushConsumer.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("获取到值: %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Println("读取消息失败")
	}

	_ = pushConsumer.Start()

	time.Sleep(time.Hour)

	_ = pushConsumer.Shutdown()
}
