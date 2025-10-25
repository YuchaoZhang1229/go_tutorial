## Gin 重定向详解
### 1. 重定向的基本概念
在 Web 开发中，重定向是指服务器接收到客户端请求后，返回一个包含新 URL 的特定状态码，引导客户端自动向新 URL 发起新请求的过程。在 Gin 框架中，重定向主要通过 `c.Redirect()` 方法实现

### 2. Gin 中的重定向方法
#### 基本重定向语法
在 Gin 中，你可以使用 c.Redirect(code int, location string)方法实现重定向：
```go
r.GET("/old", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "/new")
})
```
#### HTTP重定向
外部重定向指向不同域名或完全不同的 URL：
```go
r.GET("/baidu", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
})
```
#### 路由内部重定向
内部重定向是指在同一应用内将请求转发到不同的路由。Gin 提供了 HandleContext方法实现内部跳转
```go
r.GET("/a", func(c *gin.Context) {
    // 跳转到/b对应的路由处理函数
    c.Request.URL.Path = "/b"
    r.HandleContext(c)
})

r.GET("/b", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "b"})
})
```

### 3. HTTP 重定向状态码
Gin 支持多种 HTTP 重定向状态码，正确选择状态码对 SEO 和浏览器行为有重要影响

| 状态码  | 常量                           | 含义      | 适用场景               |
|:-----|:-----------------------------|:--------|:-------------------|
| 301  | http.StatusMovedPermanently  | 永久重定向   | 域名变更、网站改版后旧链接永久指向新链接|
| 302  | http.StatusFound             | 临时重定向   | 临时维护页面、登录后跳转等临时性重定向|
| 307  | http.StatusTemporaryRedirect | 临时重定向   | 与302类似，但保证请求方法和体不会改变|
| 308  | http.StatusPermanentRedirect | 永久重定向   | 与301类似，但保证请求方法和体不会改变|

### 4. 常见应用场景
- **用户认证与权限控制**：当用户访问需要登录的资源而未认证时，**临时重定向 (302)** 到登录页面。登录成功后，再重定向回用户最初请求的页面
- ......