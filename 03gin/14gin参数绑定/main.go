package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"` // 可从form或JSON中获取username，且必填
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 获取querystring参数
	// http://127.0.0.1:8080/form?username=sb&password=123456
	r.GET("/user", func(c *gin.Context) {
		var user UserInfo // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	// 获取form参数
	// 用Postman调试, 改成POST请求->Body->form-data->填写key-value
	// http://127.0.0.1:8080/form
	r.POST("/form", func(c *gin.Context) {
		var user UserInfo // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	// 获取json参数
	// 用Postman调试, 改成POST请求->Body->raw->选择JSON格式
	// http://127.0.0.1:8080/form
	r.POST("/json", func(c *gin.Context) {
		var user UserInfo // 声明一个UserInfo类型的变量u
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":8080")
}
