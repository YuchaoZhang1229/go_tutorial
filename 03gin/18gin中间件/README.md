## Gin中间件详解

Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些**公共的业务逻辑**，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
- **登录认证**：查看是否登录，不登录就跳转登录页面
- **权限校验**：查看是否是VIP，不是的话就跳转到VIP购买页面
- **数据分页**：一个页面显示不了，分好几页来返回
- **记录日志**：记录一些IP访问信息，放问频率等等
- **耗时统计**：统计耗时

### 1. 定义中间件
**功能**: 创建一个可复用的处理函数，用于在请求前后执行特定逻辑（如认证、日志）, 中间件必须是 gin.HandlerFunc 类型，即一个接收 *gin.Context 参数的函数

**语法**:
```go
func MiddlewareName() gin.HandlerFunc {
return func(c *gin.Context) {
    // 中间件逻辑
    // 可调用 c.Next() 或 c.Abort()
  }
}
```

### 2. 注册中间件
**功能**：将中间件应用到全局、路由组或单个路由

**语法**:
- 全局注册 (所有路由都会生效):
```go
router.Use(Middleware1(), Middleware2())
```
- 路由组注册:
```go
group := router.Group("/admin", Middleware1())
```
或
```go
group := router.Group("/admin")
group.Use(Middleware1())
```

- 单个路由注册:
```go
router.GET("/path", Middleware1(), Middleware2(), finalHandler)
```


### 3. 默认中间件
**功能**：Gin 提供了一些开箱即用的常用中间件。

**语法**：
```go
router := gin.Default()// 默认包含 Logger和 Recovery

// 或

router := gin.New(); 
router.Use(gin.Logger(), gin.Recovery())
```

### 4. c.Next() 和 c.Abort
**功能**：控制中间件链的执行流程。

**语法**：
```go
c.Next(): // 调用后续中间件和处理函数
c.Abort(): // 终止当前中间件链，阻止后续处理
c.AbortWithStatus(code int): // 中断并设置状态码
c.AbortWithStatusJSON(code, gin.H{"error": msg}): // 中断并返回 JSON 响应
```

### 5. c.Set() 和 c.Get()
**功能**：在中间件与处理函数间或不同中间件间传递数据。

**语法**：
````go
c.Set("userID", userID)             //存储值
value, exists := c.Get("userID") // 检索值 
````
