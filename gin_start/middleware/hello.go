package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	router := gin.New()

	// 使用logger和recovery 全局所有
	router.Use(gin.Logger(), gin.Recovery())

	router.Use(MyLogger())

	// 为单个路由添加中间件
	/*auth := router.Group("/goods")
	auth.Use(gin.Logger())
	auth.Use(MyLogger())*/

	router.GET("/test", MyLogger(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})

	})

	router.Run(":8080")

}

func MyLogger() gin.HandlerFunc {

	return func(c *gin.Context) {

		start := time.Now()

		//我们前端传过来的token，需要解析之后，交给具体的函数使用。
		c.Set("example", "12345")

		// 让原本该执行的逻辑继续执行
		c.Next()

		end := time.Since(start)

		fmt.Printf("请求结束，耗时：%V\n", end)

		status := c.Writer.Status() // 获取请求状态码
		fmt.Println("状态：", status)
	}

}
