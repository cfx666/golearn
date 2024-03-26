package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.Use(MyToken())

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":8080")

}

func MyToken() gin.HandlerFunc {

	return func(c *gin.Context) {

		var token string
		// 获取前端请求头里面的token。注意前端设置的请求头为x_token，我们使用gin来获取的时候，需要首字母大写X_Token
		// 取到之后，v是一个切片，其实只有一个值也是切片类型，使用v[0]取出来
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
		}

		if token != "tom" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token error"})
			//return
			c.Abort()
		}

		c.Next()
	}
}
