package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var total int32
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	/*go func() {
		fmt.Println("Hello, World!")
	}()
	time.Sleep(2 * time.Second)*/

	/*for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	//解决办法1
	for i := 0; i < 100; i++ {
		temp := i
		go func() {
			fmt.Println(temp)
		}()
	}*/

	//解决办法2
	/*var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1) //每启动一个goroutine，wg+1
		go func(i int) {
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("main done")*/

	/*wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)*/

	//var num int
	/*var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)

	go func() {
		time.Sleep(3 * time.Second)
		defer wg.Done()
		rwlock.Lock() //写锁，写锁是排他锁，读锁是共享锁
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		fmt.Println("读锁被阻塞")
		time.Sleep(10 * time.Second)
	}()

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock() //读锁，读锁是共享锁
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}
		}()
	}

	wg.Wait()*/

	var msg chan string
	msg = make(chan string, 2)

	go func(msg chan string) {

		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
	}(msg)

	/*msg <- "hello"
	msg <- "world"

	close(msg)
	msg <- "!"    //放值到channel中，但是channel已经关闭了，会报错
	data := <-msg //已经关闭的channel可以继续取值，但是不能再放值了

	time.Sleep(10 * time.Second)*/

	/*var chan1 chan int   //双向channel
	var chan2 chan<- int //只写channel
	var chan3 <-chan int //只读channel*/

	//双向channel可以转换为单向channel，但是单向channel不能转换为双向channel
	/*	chan1 = make(chan int, 2)
		var send chan<- int = chan1
		var recv <-chan int = chan1

		send <- 1*/
	//recv <- 2 //只读channel不能写入值
	//<-recv //可以不加变量接收
	//<- send //只写channel不能读入值

	/*c := make(chan int)
	go produce(c)
	go consume(c)
	*/
	/*go printNum()
	go printLetter()
	num <- true
	time.Sleep(10 * time.Second)*/

	//var ch1 = make(chan struct{})
	//var new_helloword = make(chan struct{})
	//go g1(ch1)
	//go g2(new_helloword)
	//
	//var timer = time.NewTimer(5 * time.Second)
	//for {
	//
	//	select {
	//	case <-ch1:
	//		fmt.Println("g1 done")
	//	case <-new_helloword:
	//		fmt.Println("g2 done")
	//	case <-timer.C:
	//		fmt.Println("outtime")
	//		return
	//	}
	//}

	var ctx, _ = context.WithTimeout(context.Background(), 6*time.Second)
	ctxValue := context.WithValue(ctx, "traceid", "123456")
	wgcpu.Add(1)
	go printCPUInfo(ctxValue)
	wgcpu.Wait()
	fmt.Println("监控完成")
}

var wgcpu sync.WaitGroup

func printCPUInfo(ctx context.Context) {
	defer wgcpu.Done()

	//记录一些日志，这次请求是哪个traceid发起的
	fmt.Println(ctx.Value("traceid"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出监控")
			return
		default:
			fmt.Println("cpu信息")
			time.Sleep(2 * time.Second)
		}
	}
}

var flag = make(chan struct{})

func g1(ch chan struct{}) {
	fmt.Println("g1开始运行")
	time.Sleep(1 * time.Second)
	fmt.Println("g1运行结束")

	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	fmt.Println("g2开始运行")
	time.Sleep(1 * time.Second)
	fmt.Println("g2运行结束")

	ch <- struct{}{} //g2执行完毕，向flag发送数据。这里不需要加锁，因为flag是一个channel，是线程安全的
}

var num, letter = make(chan bool), make(chan bool)

func printNum() {

	var i = 1
	for {
		//判断num是否有数据
		<-num

		fmt.Printf("%d%d", i, i+1)
		i += 2

		//通知letter可以打印了
		letter <- true
	}
}

func printLetter() {
	var str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var i = 0
	for {
		if i >= len(str) {
			return
		}

		//判断letter是否有数据
		<-letter

		fmt.Print(str[i : i+2])
		i += 2

		//通知num可以打印了
		num <- true
	}
}

func produce(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consume(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func add() {
	defer wg.Done()
	for i := 0; i < 100000000; i++ {
		/*lock.Lock()
		total += 1
		lock.Unlock()*/
		atomic.AddInt32(&total, 1)
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000000; i++ {
		/*lock.Lock()
		total -= 1
		lock.Unlock()*/
		atomic.AddInt32(&total, 1)
	}
}

func asyncPrint() {
	fmt.Println("Hello, World!")
}
