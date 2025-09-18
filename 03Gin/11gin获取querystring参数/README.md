## gin框架获取querystring参数
- GET请求 URL ？后面的是querystring参数
- key=value格式，多个key-value用 & 连接
- http://127.0.0.1:8080/web?query=小王子&age=18
- 有多种方法
  - key = c.Query("key")
  - key = c.DefaultQuery("key", "defaultValue")
  - key, ok = c.GetQuery("key")