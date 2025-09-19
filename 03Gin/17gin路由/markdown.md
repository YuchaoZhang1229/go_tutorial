## Gin路由详解
### 普通路由
RESTful API 支持标准 HTTP 方法：
- 查询：r.GET("/path", handler)
- 新增：r.POST("/path", handler)
- 更新：r.PUT("/path", handler)
- 删除：r.DELETE("/path", handler)
- 任意: r.Any("/path", handler)
- 无路由: r.NoRoute(handler)


### 路由组
我们可以将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别。
```go
func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {...})
		userGroup.GET("/login", func(c *gin.Context) {...})
		userGroup.POST("/login", func(c *gin.Context) {...})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
	}
	r.Run()
}
```
路由组也是支持嵌套的，例如：
```go
shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {...})
		shopGroup.GET("/cart", func(c *gin.Context) {...})
		shopGroup.POST("/checkout", func(c *gin.Context) {...})
		// 嵌套路由组
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {...})
	}
```
通常我们将路由分组用在划分业务逻辑或划分API版本时。

### 路由原理
Gin框架中的路由使用的是httprouter这个库。

其基本原理就是构造一个路由地址的**前缀树**。

