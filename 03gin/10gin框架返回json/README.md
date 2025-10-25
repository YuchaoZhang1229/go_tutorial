gin框架返回json有两种方法
1. 使用map
```go
// gin.H等价于map[string]interface{}
data := gin.H{
    "name":    "小王子",
    "message": "Hello World!",
    "age":     18,
}

```
2. 使用结构体

结构体的成员名字首字母名字需要大写

如果一定要小写那么就在对应字段后面加tag
```go
type msg struct {
    Name    string `json:"name"`
    Message string
    Age     int
}

data := msg{
Name:    "小王子",
Message: "Hello Golang!",
Age:     18,
}
```