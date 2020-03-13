package main

import (
	"fmt"
	"github.com/gin-gonic/gin/testdata/protoexample"
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
	r.POST("/login_json", loginJson)
	r.POST("/login_form", loginForm)
	r.GET("/:user/:password", loginURI)

	// 多种响应方式
	// 1.json
	r.GET("/resp_json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "somejson", "status": 200})
	})

	// 2.结构体
	r.GET("/resp_struct", func(ctx *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}

		msg.Message = "hello"
		msg.Name = "world"
		msg.Number = 10086
		ctx.JSON(200, msg)
	})

	// 3.XML
	r.GET("/resp_xml", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{"message": "abcd"})
	})

	// 4.YAML
	r.GET("/resp_yaml", func(ctx *gin.Context) {
		ctx.YAML(200, gin.H{"name": "cdd"})
	})

	// 5 protobuf
	r.GET("/resp_protobuf", func(ctx *gin.Context) {
		reps := []int64{0, 1, 2}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		ctx.ProtoBuf(200, data)
	})

	r.Run(":9000")
}
