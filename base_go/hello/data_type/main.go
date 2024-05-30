package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(ch chan<- int) { // 生产，只写。只要该通道具有写能力就行
	for {
		ch <- rand.Intn(10)
		time.Sleep(1 * time.Second)
	}
}
func consume(ch <-chan int) { // 消费，只读。只要该通道具有读能力就行
	for {
		t := <-ch
		fmt.Println("消费，从只读通道接收", t)
	}
}
func main() {

	count := make(chan int, 4)
	fin := make(chan bool)
	newBase := 1000
	t1 := time.NewTicker(time.Second)
	t2 := time.NewTicker(5 * time.Second)
	go func() {
		defer func() { fin <- true }()
		for i := 0; i < 4; i++ {
			count <- i
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(len(count), "~~~~@@@")
	for {
		select { // 监听多路通道
		case <-t1.C:
			fmt.Println("每隔一秒看看长度", len(count))
		case <-t2.C:
			fmt.Println("每隔5秒取一次", <-count)
		case count <- newBase: // 发送数据成功进入通道执行该case
			newBase++
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
	}
}
