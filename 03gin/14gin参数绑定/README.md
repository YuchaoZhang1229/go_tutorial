## gin框架参数绑定
- 将请求参数自动绑定到结构体
- 主要获取方法:
    - key = c.ShouldBind(&struct{})
- 常用场景: 简化参数获取和验证，处理多种格式的请求数据

## c.ShouldBind(&struct{})代码执行理解

Gin 框架中的 c.ShouldBind(&user)是一个非常方便的方法，它能帮你自动将 HTTP 请求中的数据解析并绑定到你指定的 Go 结构体实例上。它会根据请求的 Content-Type头部信息**自动选择**合适的绑定方式。

下面是你提供的 GET 接口代码中 c.ShouldBind(&user)的**简要执行逻辑**：
1. **接收请求与自动选择绑定器**：当 GET 请求 /user?username=abc&password=123到来时，ShouldBind会根据请求方法和 Content-Type（对于 GET 请求，Content-Type 通常不重要）自动选择绑定器。由于是 GET 请求，它会使用 formBinding或 queryBinding来处理 URL 查询参数（即 ?后面的部分）
2. **映射字段**：ShouldBind通过**反射**检查 user变量（类型为 UserInfo）的结构体字段标签（Tag）。它尤其会查找 form标签（例如 form:"username"）
3. **提取参数与赋值**：根据结构体字段的 form 标签名（如 username），ShouldBind会从请求的**查询参数**中寻找同名的参数值，并将其转换后填充到 UserInfo结构体的对应字段中
4. **返回结果**
   - **成功**：若所有参数均正确映射，err为 nil，user变量被正确赋值，随后打印结构体内容并返回 HTTP 200 状态码及 "ok" 状态
   - **失败**：出现错误（如必填字段为空或类型转换失败），err不为 nil，返回 HTTP 400 状态码及错误信息