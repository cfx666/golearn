package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"

	"net/http"
)

type Login struct {
	User     string `json:"user" binding:"required,min=5,max=10"`
	Password string `json:"password" binding:"required"`
}

type Sign struct {
	Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // eqfield=Password表示这个字段的值必须和Password字段的值相等
}

var trans ut.Translator

func main() {

	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化翻译器错误")
		return
	}

	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBind(&json); err != nil {
			err, ok := err.(validator.ValidationErrors)
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": removeTopStruct(err.Translate(trans))})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "登陆成功"})
	})

	router.POST("/signJSON", func(c *gin.Context) {
		var json Sign
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
	})

	router.Run(":8080")

}

func InitTrans(locale string) (err error) {
	//	修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("json")
			name = strings.SplitN(name, ",", 2)[0]
			if name == "-" { // -是json里面的约束，表示不需要处理
				return ""
			}
			return name

		})

		zh := zh.New() // 中文翻译器
		en := en.New() // 英文翻译器
		// 第一个参数是备用（英文）翻译器，后面的参数是应该支持的语言环境
		uni := ut.New(en, zh, zh)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			err = en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for key, value := range fileds {
		rsp[key[strings.Index(key, ".")+1:]] = value
	}
	return rsp

}
