package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	MysqlInfo  MysqlConfig `mapstructure:"mysql"`
}

func main() {

	// 1.创建一个viper对象
	v := viper.New()
	// 2.设置配置文件的名字
	v.SetConfigFile("viper_test/ch01/config.yaml")
	// 3.添加配置文件所在的路径
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//fmt.Println(v.Get("name"))

	var serverConfig ServerConfig
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
}
