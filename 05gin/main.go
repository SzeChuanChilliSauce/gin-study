package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
type Login struct {
	// binding:"required"修饰的字段是必需字段，若接收为空，报错
	User     string `form:"user" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password"`
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

	ctx.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}

func main() {
	r := gin.Default()
	r.POST("/loginjson", loginJson)
	r.Run(":9000")
}
