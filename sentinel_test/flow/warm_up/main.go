package main

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
	"math/rand"
	"time"
)

func main() {

	// 初始化sentinel
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("初始化sentinel失败：%v", err)
	}

	// 配置规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "test",      // 资源名称，这个是限流规则的唯一标识
			TokenCalculateStrategy: flow.WarmUp, // 限流的计算策略
			ControlBehavior:        flow.Reject, // 流量控制的行为，Reject：拒绝
			Threshold:              1000,        // 阈值
			WarmUpPeriodSec:        30,          // 预热时间，在这个时间内，限流阈值从0逐渐增加到设定值
			//WarmUpColdFactor:       3,           // 预热因子
		},
	})
	if err != nil {
		log.Fatalf("加载规则失败：%v", err)
	}

	var globalTotal int
	var blockTotal int
	var passTotal int

	var ch = make(chan struct{})

	// 每一秒统计一次 这一秒通过了多少 block了多少
	for i := 0; i < 100; i++ {
		go func() {
			for {
				globalTotal++
				e, b := sentinel.Entry("test", sentinel.WithTrafficType(base.Inbound))
				if b != nil {
					blockTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					passTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					e.Exit()
				}
			}
		}()
	}

	// 每秒统计一次
	go func() {
		var oldTotal int // 过去一秒的总数
		var oldBlock int // 过去一秒的阻塞数
		var oldPass int  // 过去一秒的通过数
		for {
			oneSecond := globalTotal - oldTotal
			oldTotal = globalTotal

			oneSecondBlock := blockTotal - oldBlock
			oldBlock = blockTotal

			oneSecondPass := passTotal - oldPass
			oldPass = passTotal

			time.Sleep(time.Second)
			log.Printf("全局:%d, 通过:%d, 阻塞:%d", oneSecond, oneSecondPass, oneSecondBlock)
		}
	}()

	<-ch
}
