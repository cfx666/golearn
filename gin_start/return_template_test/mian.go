package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	router := gin.Default()

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("dir: %s\n", dir) //C:\Users\xy\AppData\Local\JetBrains\IntelliJIdea2023.1\tmp\GoLand

	// LoadHTMLFiles函数将指定目录下的文件加载好
	//router.LoadHTMLFiles("D:\\workspace\\GoProject\\learngo\\gin_start\\template_test\\templates\\index.tmpl")
	//router.LoadHTMLFiles("templates/index.tmpl")
	//router.LoadHTMLFiles("templates/index.tmpl", "templates/books.html")
	//router.LoadHTMLGlob("templates/*")
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "测试返回html模板",
		})
	})

	router.GET("/books", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books/list.html", gin.H{
			"name": "微服务实战",
		})
	})

	router.GET("/users/userlist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"user": "用户列表",
		})
	})

	router.GET("/goods/goodslist", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", gin.H{
			"goods": "商品列表",
		})
	})

	router.Run(":8080")
}
