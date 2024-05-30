package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"learngo/nacos_test/config"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         "9bd04558-3449-4831-839a-ce0784ad5818", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev"})
	//fmt.Println(content)
	SrvConfig := config.ServerConfig{}
	err = json.Unmarshal([]byte(content), &SrvConfig)
	fmt.Println(SrvConfig)

	/*	_ = client.ListenConfig(vo.ConfigParam{
			DataId: "user-web.yaml",
			Group:  "dev",
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			},
		})

		time.Sleep(3000 * time.Second)*/
}
