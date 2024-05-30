package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
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
			TokenCalculateStrategy: flow.Direct, // 限流的计算策略，Direct：直接计数，Threshold的值就是次数
			ControlBehavior:        flow.Reject, // 流量控制的行为，Reject：拒绝
			Threshold:              10,          // 阈值
			StatIntervalInMs:       1000,        // 统计时间窗口，单位毫秒
		},
	})
	if err != nil {
		log.Fatalf("加载规则失败：%v", err)
	}

	for i := 0; i < 12; i++ {
		e, b := sentinel.Entry("test", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			fmt.Println("限流了")
		} else {
			fmt.Println("通过")
			e.Exit()
		}
	}
}
