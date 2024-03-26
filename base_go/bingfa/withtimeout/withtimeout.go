package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var ctx, _ = context.WithTimeout(context.Background(), 6*time.Second)
	wgcpu.Add(1)
	go printCPUInfo(ctx)

	wgcpu.Wait()
	fmt.Println("监控完成")
}

var wgcpu sync.WaitGroup

func printCPUInfo(ctx context.Context) {
	defer wgcpu.Done()
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
