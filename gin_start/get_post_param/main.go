package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/get", Get)
	r.POST("/form_post", formPost)
	r.POST("/both", getPost)

	r.Run(":8089")
}

func getPost(c *gin.Context) {

	name := c.DefaultQuery("name", "get_default_name")
	age := c.DefaultQuery("age", "get_default_age")

	nick := c.DefaultPostForm("nick", "post_default_nick")
	message := c.DefaultPostForm("message", "post_default_message")

	c.JSON(http.StatusOK, gin.H{
		"nick":    nick,
		"message": message,
		"name":    name,
		"age":     age,
	})

}

func formPost(c *gin.Context) {
	fmt.Printf("formPost\n")
	nick := c.DefaultPostForm("nick", "post_default_nick")
	message := c.DefaultPostForm("message", "post_default_message")

	//如果不希望有默认值，可以这样写
	//nick := c.PostForm("nick")

	c.JSON(http.StatusOK, gin.H{
		"nick":    nick,
		"message": message,
	})
}

func Get(c *gin.Context) {
	name := c.DefaultQuery("name", "get_default_name")
	age := c.DefaultQuery("age", "get_default_age")

	//如果不希望有默认值，可以这样写
	//name := c.Query("name")       //是c.Request.URL.Query().Get("lastname") 的简写

	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
