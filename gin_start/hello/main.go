package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func pong(c *gin.Context) {

	// gin.h点进去看是一个map[string]interface{}类型。所以这里的第二个参数是一个map，可替换为map
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	// 实例化一个默认的server
	r := gin.Default()
	// get方法，第一个参数是路径，第二个参数是处理函数，点进去get方法可以看到第二个方法是一个HandlerFunc类型，且是可变参数
	// HandlerFunc是一个函数类型，匿名的。它的定义是：type HandlerFunc func(*Context)
	r.GET("/ping", pong)
	r.Run() // listen and serve on 0.0.0.0:8080
}
