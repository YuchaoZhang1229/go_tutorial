## gin框架获取path参数
- 请求的参数通过URL路径传递, 如 /:name/:age 中的 name 和 age
- 访问http://127.0.0.1:8080/小王子/18发送GET请求, 
- 主要获取方法:
    - key = c.Param("key")
- 常用场景: `用户登录`