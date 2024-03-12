package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

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
