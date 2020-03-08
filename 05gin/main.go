package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
type Login struct {
	// binding:"required"修饰的字段是必需字段，若接收为空，报错
	User     string `form:"user" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// json数据的解析和绑定
func loginJson(ctx *gin.Context) {
	var loginData Login

	// 将request的body中的数据自动按照json格式解析到结构体
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断用户名密码是否正确
	if !(loginData.User == "cdd" && loginData.Password == "cdd123") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "注册成功"})
}

// 表单数据解析和绑定
func loginForm(ctx *gin.Context) {
	var formData Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := ctx.Bind(&formData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !(formData.User == "cdd" && formData.Password == "cdd123") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "注册成功"})
}

// URI数据的解析与绑定
func loginURI(ctx *gin.Context) {
	var uriData Login

	if err := ctx.ShouldBindUri(&uriData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(uriData.User, uriData.Password)
	fmt.Println(ctx.Param("password"))

	// ctx.Param("user")
	// ctx.Param("password")

	if !(uriData.User == "cdd" && uriData.Password == "cdd123") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名或密码错误"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "msg": "注册成功"})
}

func main() {
	r := gin.Default()
	r.POST("/loginjson", loginJson)
	r.POST("/loginform", loginForm)
	r.GET("/:user/:password", loginURI)
	r.Run(":9000")
}
