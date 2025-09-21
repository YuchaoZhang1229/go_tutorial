package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// HandlerFunc
func homeHandler(c *gin.Context) {
	fmt.Println("home in ...")
	name, ok := c.Get("name") // 从上下问中取值,可以实现跨中间件取值
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
		"name":    name,
	})
}

// 定义一个中间件m1: 统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()
	c.Next() // 调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost:", cost)
	fmt.Println("m1 out ...")
}

// 定义一个中间件m2:
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	//c.Abort() // 阻止调用后续的处理函数
	c.Set("name", "q1mi") // 在上下文c中设置值
	fmt.Println("m2 out ...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 或者一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			// 是否登录
			// if 是登录用户
			// c.Next()
			// else
			// c.Abort()
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(m1, m2, authMiddleware(true)) // 全局注册中间件
	r.GET("/home", homeHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})

	// 路由组注册中间件方法1:
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "xx/home",
			})
		})
	}

	// 路由组注册中间件方法2:
	ooGroup := r.Group("/oo")
	ooGroup.Use(authMiddleware(true))
	{
		ooGroup.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "oo/home",
			})
		})
	}
	r.Run(":8080")

}
