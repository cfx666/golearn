package main

import (
	"github.com/gin-gonic/gin"
	"learngo/gin_start/return_json_protobuf/proto"
	"net/http"
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.GET("/json", json)
	r.GET("/protobuf", protobuf)

	r.GET("/json_html", Json_html)
	r.GET("/pure_json", pureJson)

	r.Run(":8089")
}

func json(c *gin.Context) {

	user := struct {
		Name string `json:"name"`
		Age  int
	}{
		Name: "zhangsan",
		Age:  18,
	}

	c.JSON(http.StatusOK, user)

}

func protobuf(c *gin.Context) {

	var course = []string{"java", "c++", "python"}

	teacher1 := &proto.Teacher{
		Name:   "zhangsan",
		Course: course,
	}

	c.ProtoBuf(http.StatusOK, teacher1)

}

func Json_html(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"html": "<b>hello world</b>",
	})

}

func pureJson(c *gin.Context) {

	c.PureJSON(http.StatusOK, gin.H{
		"html": "<b>hello world</b>",
	})

}
