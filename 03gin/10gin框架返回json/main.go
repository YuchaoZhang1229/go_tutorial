package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		// 方法1: 使用map
		//data := map[string]interface{}{
		//	"name":    "小王子",
		//	"message": "Hello World!",
		//	"age":     18,
		//}
		data := gin.H{
			"name":    "小王子",
			"message": "Hello World!",
			"age":     18,
		}
		c.JSON(http.StatusOK, data)
	})

	// 方法2: 结构体. 灵活使用tag来对结构体字段做定制化操作
	type msg struct {
		Name    string `json:"name"`
		Message string
		Age     int
	}

	r.GET("/json2", func(c *gin.Context) {

		data := msg{
			Name:    "小王子",
			Message: "Hello Golang!",
			Age:     18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":8080")
}
