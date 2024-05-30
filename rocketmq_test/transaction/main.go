package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"time"
)

type OrderListener struct{}

func (t *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	// 模拟订单业务的处理
	fmt.Println("模拟订单业务的处理")
	time.Sleep(time.Second * 3)

	//fmt.Println("模拟订单业务处理成功")
	//return primitive.CommitMessageState

	//fmt.Println("模拟订单业务处理失败")
	//return primitive.RollbackMessageState

	// 模拟未知状态，本地事务处理逻辑无缘无故失败。返回未知状态，rocketmq会自己回查
	return primitive.UnknowState
}

func (t *OrderListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("rocketmq的消息回查")
	// 自己看业务逻辑，判断本地事务是否成功。就算回查出现异常，下次运行不会丢失消息，还是会回查
	time.Sleep(time.Second * 15)
	return primitive.CommitMessageState
}

func main() {
	transactionProducer, _ := rocketmq.NewTransactionProducer(&OrderListener{}, producer.WithNameServer([]string{"192.168.137.129:9876"}))

	_ = transactionProducer.Start()

	res, err := transactionProducer.SendMessageInTransaction(context.Background(), &primitive.Message{
		Topic: "TransTopic",
		Body:  []byte("Hello,by transaction message!"),
	})
	if err != nil {
		fmt.Printf("发送失败: %s\n", err)
	} else {
		fmt.Printf("发送成功: result=%s\n", res.String())
	}

	// 测试回查，这里等待一段时间，模拟回查
	time.Sleep(time.Hour)

	if err = transactionProducer.Shutdown(); err != nil {
		fmt.Printf("shutdown producer error: %s", err)
	}

}
