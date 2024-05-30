package main

import "fmt"

/*
变量

	声明和初始化
		1.单声明
		2.多声明
	分类
		1.函数内变量
		2.包内变量
	注意
	作用域
	匿名变量
*/
func main() {

	// 单声明
	var a int
	var b string = "abc"
	var c = "def"
	d := 1

	fmt.Println(a, b, c, d)

	// 多声明
	/*	var l,n int
		var f, g string = "abc", "def"
		var e, t = "def", 1
		k, j := 1, false
		var (
			q int
			w string = "abc"
			m        = "def"
		)*/
}
