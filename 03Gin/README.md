Go世界里最流行的Web框架，[Github](https://github.com/gin-gonic/gin)上有84K+star。 基于[httprouter](https://github.com/julienschmidt/httprouter)开发的Web框架。 [中文文档](https://gin-gonic.com/zh-cn/docs/)齐全，简单易用的轻量级框架。

```go
// 项目初始化
go mod init
// 安装gin
go get -u github.com/gin-gonic/gin
// 运行
go run main.go
```

## Web框架需要做什么

我们先思考下，一个完整的Web开发框架需要做哪些事情

| 组件            | 功能                                                         | 是否必须 |
| --------------- | ------------------------------------------------------------ | -------- |
| server          | 作为server，监听端口，接受请求                               | 是       |
| router          | 路由和分组路由，可以把请求路由到对应的处理函数               | 是       |
| middleware      | 支持中间件，对外部发过来的http请求经过中间件处理，再给到对应的处理函数。例如http请求的日志记录、请求鉴权(比如校验token)、CORS支持、CSRF校验等。 | 是       |
| template engine | 模板引擎，支持后端代码对html模板里的内容做渲染(render)，返回给前端渲染好的html | 否       |
| ORM             | 对象关系映射，可以把代码里的对象和关系数据库的表、字段做映射关联，通过操作对象来实现数据库的增删查改等操作。 | 否       |


## Gin有什么

Gin的主要作者是[Manu](https://github.com/manucorporat)，[Javier](https://github.com/javierprovecho)和[Bo-Yi Wu](https://github.com/appleboy)，2016年发布第一个版本，目前是最受欢迎的开源Go框架。

Gin除了支持上面表格里列的server、router、middleware和template之外，还支持

* Crash-free：崩溃恢复，Gin可以捕捉运行期处理http请求过程中的panic并且做recover操作，让服务一直可用。
* JSON validation：JSON验证。Gin可以解析和验证request里的JSON内容，比如字段必填等。当然开发人员也可以选择使用第三方的JSON validation工具，比如[beego validation](https://github.com/beego/beego/tree/develop/core/validation)。
* Error management：错误管理。Gin提供了一种简单的方式可以收集http request处理过程中的错误，最终中间件可以选择把这些错误写入到log文件、数据库或者发送到其它系统。
* Middleware Extendtable：可以自定义中间件。Gin除了自带的官方中间件之外，还支持用户自定义中间件，甚至可以把自己开发的中间件提交到[官方代码仓库](https://github.com/gin-gonic/contrib)里。

Gin本身不支持ORM，如果想在Gin框架里使用ORM，可以选择使用第三方的ORM，比如[gorm](https://github.com/go-gorm/gorm)。