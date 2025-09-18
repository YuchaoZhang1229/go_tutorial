package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器那边发请求携带的 query string 参数
		name := c.Query("query")
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody")
		//name, ok := c.GetQuery("query")
		//if !ok {
		//	// 取不到
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":8080")
}
