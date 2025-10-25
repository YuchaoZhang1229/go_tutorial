package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询：r.GET("/path", handler)
// 新增：r.POST("/path", handler)
// 更新：r.PUT("/path", handler)
// 删除：r.DELETE("/path", handler)

func main() {
	r := gin.Default()

	// 访问/index的GET请求会走这一条处理逻辑
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.Any("/home", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		// ...
		default:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		}
	})

	// NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "NoRoute",
		})
	})

	// 视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "/video/index",
	//	})
	//}
	//
	//r.GET("/video/xx", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "/video/xx",
	//	})
	//})
	//
	//r.GET("/video/oo", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "/video/oo",
	//	})
	//})

	// 路由组
	// 把共用的前缀提取出来, 创建一个路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/index",
			})
		})

		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/xx",
			})
		})

		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/oo",
			})
		})
	}

	// 商城的首页和详情页
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "/shop/index",
		})
	})

	r.GET("/shop/xx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "/shop/xx",
		})
	})

	r.GET("/shop/oo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "/shop/oo",
		})
	})

	r.Run(":8080")
}
