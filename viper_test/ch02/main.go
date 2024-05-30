package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string      `mapstructure:"name"`
	MysqlInfo  MysqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main() {

	/*debug := GetEnvInfo("MXSHOP_DEBUG")
	var configFilePrefix = "config"
	var configFileName = fmt.Sprintf("viper_test/ch02/%s-pro.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("viper_test/ch02/%s-debug.yaml", configFilePrefix)
	}*/

	// 1.创建一个viper对象
	v := viper.New()
	// 2.设置配置文件的名字
	v.SetConfigFile("viper_test/ch02/config-debug.yaml")
	// 3.添加配置文件所在的路径
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//fmt.Println(v.Get("name"))

	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	// viper动态监控变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file change", e.Name)

		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(&serverConfig)
	})
	v.WatchConfig()

	time.Sleep(300 * time.Second)
}
