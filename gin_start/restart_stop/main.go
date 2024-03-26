package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 优雅退出，当我们关闭程序的时候应该做的后续退出
	// 微服务 启动之前 或者启动之后会做一件事：将当前的服务ip和端口号注册到注册中心
	// 当我们关闭服务的时候，需要将当前的服务从注册中心移除，也就是需要通知注册中心

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})

	go func() {
		router.Run(":8080")
	}()

	// 我们把router.Run(":8080")放到goroutine中，这样代码会继续向下执行，我们使用一个chan，当收到系统退出的信号时，那么chan不会阻塞

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 执行退出的操作
	fmt.Println("关闭服务...")
	fmt.Println("通知注册中心，注销服务...")
}
