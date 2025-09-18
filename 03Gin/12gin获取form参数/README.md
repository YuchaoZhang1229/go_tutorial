## gin框架获取form参数
- 提交表单发送POST请求, key对应的是<input name='xxx'></input>中的name的值
- 访问http://127.0.0.1:8080/login发送GET请求, 提交表格的时候发送的POST请求
- 主要获取方法:
  - key = c.PostForm("key")
  - key = c.DefaultPostForm("key", "default_value")
  - key, ok = c.GetPostForm("key")
- 常用场景: `用户登录`, `用户注册`,`用户提交表单`