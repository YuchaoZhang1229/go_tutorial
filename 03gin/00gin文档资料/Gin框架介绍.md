

# Gin框架介绍

### 介绍

Gin 是一个用 Go (Golang) 编写的 Web 框架。 它具有类似 martini 的 API，性能要好得多，多亏了 [httprouter](https://github.com/julienschmidt/httprouter)，速度提高了 40 倍。



### 快速入门

安装gin

```
go get -u github.com/gin-gonic/gin
```

引入gin

```
import "github.com/gin-gonic/gin"
```

开始

```
package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```



### 路由

- 普通路由

  ```
  router.GET("/", func)
  router.POST("/login", func)
  router.Any("/login", func)
  ```

  

- 路由分组

  ```
  // 简单的路由组: v1
    {
      v1 := router.Group("/v1")
      v1.POST("/login", loginEndpoint)
      v1.POST("/submit", submitEndpoint)
      v1.POST("/read", readEndpoint)
    }
  
    // 简单的路由组: v2
    {
      v2 := router.Group("/v2")
      v2.POST("/login", loginEndpoint)
      v2.POST("/submit", submitEndpoint)
      v2.POST("/read", readEndpoint)
    }
  ```

  

- RESTFUL

  ```
  router.GET("/user", QueryFunc)	//查询
  router.Post("/user", AddFunc)	// 新增
  router.Delete("/user", DeleteFunc)	// 删除
  router.PUT("/user", UpdateFunc)	// 更新（客户端提供完整数据）
  router.PATCH("/user", PatchUpdateFunc)	// 更新（客户端提供需要修改的数据）
  ```



- 重定向

  ```
  // 重定向到外部
  router.GET("/test", func(c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
  })
  
  // 重定向到内部
  router.POST("/test", func(c *gin.Context) {
    c.Redirect(http.StatusFound, "/foo")
  })
  
  
  router.GET("/test", func(c *gin.Context) {
      c.Request.URL.Path = "/test2"
      router.HandleContext(c)
  })
  router.GET("/test2", func(c *gin.Context) {
      c.JSON(200, gin.H{"hello": "world"})
  })
  ```

- 静态文件

  ```
  func main() {
    router := gin.Default()
    router.Static("/assets", "./assets")	// 文件目录
    router.StaticFS("/more_static", http.Dir("my_file_system"))
    router.StaticFile("/favicon.ico", "./resources/favicon.ico")	// 单独的文件
  
    // 监听并在 0.0.0.0:8080 上启动服务
    router.Run(":8080")
  }
  ```

  

### 输出

- XML/JSON/TOML/YAML/ProtoBuf

  ```
  c.JSON(http.StatusOK, struct/gin.H)
  c.XML(http.StatusOK, struct/gin.H)
  c.YAML(http.StatusOK, struct/gin.H)
  c.ProtoBuf(http.StatusOK, struct/gin.H)
  ...
  可以到定义文件中去看更多的输出方法
  ```

  

- HTML

  ```
  func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Main website",
      })
    })
    router.Run(":8080")
  }
  
  templates/index.tmpl
  <html>
    <h1>
      {{ .title }}
    </h1>
  </html>
  ```

  ```
  func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/**/*")
    router.GET("/posts/index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
        "title": "Posts",
      })
    })
    router.GET("/users/index", func(c *gin.Context) {
      c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
        "title": "Users",
      })
    })
    router.Run(":8080")
  }
  
  templates/posts/index.tmpl
  {{ define "posts/index.tmpl" }}
  <html><h1>
    {{ .title }}
  </h1>
  <p>Using posts/index.tmpl</p>
  </html>
  {{ end }}
  
  templates/users/index.tmpl
  {{ define "users/index.tmpl" }}
  <html><h1>
    {{ .title }}
  </h1>
  <p>Using users/index.tmpl</p>
  </html>
  {{ end }}
  ```

  ```
  自定义模板渲染器
  ...
  ```

  

### 参数

- 参数绑定

  ```
  ShouldBind: 自动识别参数，并绑定对应字段到结构体中
  
  ShouldBindJSON	tag:json
  ShouldBindXML		tag:xml
  ShouldBindQuery	tag:form
  ShouldBindYAML	tag:yaml
  ShouldBindTOML	tag:toml
  ShouldBindHeader	tag:header	无法自动识别
  ShouldBindUri	tag:uri	无法自动识别
  
  type formA struct {
    Foo string `json:"foo" xml:"foo" binding:"required" form:"foo"`
  }
  ```

  ```
  type formA struct {
    Foo string `json:"foo" xml:"foo" binding:"required"`
  }
  
  type formB struct {
    Bar string `json:"bar" xml:"bar" binding:"required"`
  }
  
  func SomeHandler(c *gin.Context) {
    objA := formA{}
    objB := formB{}
    // c.ShouldBind 使用了 c.Request.Body，不可重用。
    if errA := c.ShouldBind(&objA); errA == nil {
      c.String(http.StatusOK, `the body should be formA`)
    // 因为现在 c.Request.Body 是 EOF，所以这里会报错。
    } else if errB := c.ShouldBind(&objB); errB == nil {
      c.String(http.StatusOK, `the body should be formB`)
    } else {
      ...
    }
  }
  ```

  要想多次绑定，可以使用 `c.ShouldBindBodyWith`

  ```
  
  func SomeHandler(c *gin.Context) {
    objA := formA{}
    objB := formB{}
    // 读取 c.Request.Body 并将结果存入上下文。
    if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
      c.String(http.StatusOK, `the body should be formA`)
    // 这时, 复用存储在上下文中的 body。
    } else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
      c.String(http.StatusOK, `the body should be formB JSON`)
    // 可以接受其他格式
    } else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
      c.String(http.StatusOK, `the body should be formB XML`)
    } else {
      ...
    }
  }
  ```

  ```
  c.MustBindWith(&obj, binding.JSON)
  如果发生绑定错误，则请求终止，并设置错误码为400。ShouldBind()返回错误给开发者，不会导致请求终止。
  ```

  

- 非绑定获取

  ```
  r.GET("/:u/:p", func)
  c.Param("u")	// url: /11/22
  
  r.GET("/", func)
  c.Query("u")	// url: /?u=11
  c.DefaultQuery("u", "222")
  
  r.GET("/", func)
  c.PostForm("u")	// 获取表单中字段的值
  c.DefaultPostForm("u", "222")
  ```

  

### 中间件

- 调用栈

  ```
  func mw1() gin.HandlerFunc {
  	return func(c *gin.Context) {
  		fmt.Println("mw1 before")
  		c.Next()
  		fmt.Println("mw1 after")
  	}
  }
  
  func mw2() gin.HandlerFunc {
  	return func(c *gin.Context) {
  		fmt.Println("mw2 before")
  		c.Next()
  		fmt.Println("mw2 after")
  	}
  }
  
  func Middleware() {
  	r := gin.Default()
  
  	r.GET("/", mw1(), mw2(), func(c *gin.Context) {
  		fmt.Println("self")
  		c.String(http.StatusOK, "self")
  	})
  
  	err := r.Run(":8080")
  	if err != nil {
  		panic(err)
  	}
  }
  
  输出：
  mw1 before
  mw2 before
  self
  mw2 after
  mw1 after
  ```

- 用户认证

  ```
  // 模拟一些私人数据
  var secrets = gin.H{
    "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
    "austin": gin.H{"email": "austin@example.com", "phone": "666"},
    "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
  }
  
  func main() {
    router := gin.Default()
  
    // 路由组使用 gin.BasicAuth() 中间件
    // gin.Accounts 是 map[string]string 的一种快捷方式
    authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
      "foo":    "bar",
      "austin": "1234",
      "lena":   "hello2",
      "manu":   "4321",
    }))
  
    // /admin/secrets 端点
    // 触发 "localhost:8080/admin/secrets
    authorized.GET("/secrets", func(c *gin.Context) {
      // 获取用户，它是由 BasicAuth 中间件设置的
      user := c.MustGet(gin.AuthUserKey).(string)
      if secret, ok := secrets[user]; ok {
        c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
      } else {
        c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
      }
    })
  
    // 监听并在 0.0.0.0:8080 上启动服务
    router.Run(":8080")
  }使用中间件
  ```

- 使用中间件

  ```
  func main() {
    // 新建一个没有任何默认中间件的路由
    r := gin.New()
  
    // 全局中间件
    // Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
    // By default gin.DefaultWriter = os.Stdout
    router.Use(gin.Logger())
  
    // Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
    router.Use(gin.Recovery())
  
    // 你可以为每个路由添加任意数量的中间件。
    router.GET("/benchmark", MyBenchLogger(), benchEndpoint)
  
    // 认证路由组
    // authorized := router.Group("/", AuthRequired())
    // 和使用以下两行代码的效果完全一样:
    authorized := router.Group("/")
    // 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
    // AuthRequired() 中间件
    authorized.Use(AuthRequired())
    {
      authorized.POST("/login", loginEndpoint)
      authorized.POST("/submit", submitEndpoint)
      authorized.POST("/read", readEndpoint)
  
      // 嵌套路由组
      testing := authorized.Group("testing")
      testing.GET("/analytics", analyticsEndpoint)
    }
  
    // 监听并在 0.0.0.0:8080 上启动服务
    router.Run(":8080")
  }
  ```



### 模型验证

- 官方验证器

  Gin使用 [**go-playground/validator/v10**](https://github.com/go-playground/validator) 进行验证。 查看标签用法的全部[文档](https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Baked_In_Validators_and_Tags).

  ```
  type LoginInfo struct {
  	Username string `json:"username" form:"username" binding:"required"`
  	Password string `json:"password" form:"password" binding:"number"`
  	Email    string `json:"email" form:"email" binding:"email"`
  }
  
  func Validator() {
  	r := gin.Default()
  
  	r.GET("/", func(c *gin.Context) {
  		login := LoginInfo{}
  		err := c.ShouldBind(&login)
  		if err != nil {
  			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  			return
  		}
  
  		c.JSON(http.StatusOK, login)
  	})
  
  	err := r.Run(":8080")
  	if err != nil {
  		panic(err)
  	}
  }
  ```

  

- 自定义验证器

  ```
  import (
    "net/http"
    "reflect"
    "time"
  
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
  )
  
  // Booking 包含绑定和验证的数据。
  type Booking struct {
    CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
    CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
  }
  
  var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
    date, ok := fl.Field().Interface().(time.Time)
    if ok {
      today := time.Now()
      if today.After(date) {
        return false
      }
    }
    return true
  }
  
  func main() {
    route := gin.Default()
  
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
      v.RegisterValidation("bookabledate", bookableDate)
    }
  
    route.GET("/bookable", getBookable)
    route.Run(":8085")
  }
  
  func getBookable(c *gin.Context) {
    var b Booking
    if err := c.ShouldBindWith(&b, binding.Query); err == nil {
      c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
    } else {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
  }
  ```

  

### HTTPS

- 自有证书

  Openssl 生成证书，如果只是提供API服务，可以用没有经过认证的证书，如果是网站，则需要认证证书

  ```
  router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
  ```

  

- 开源免费认证证书（Let's Encrypt）

  ```
  package main
  
  import (
    "log"
  
    "github.com/gin-gonic/autotls"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/acme/autocert"
  )
  
  func main() {
    router := gin.Default()
  
    // Ping handler
    router.GET("/ping", func(c *gin.Context) {
      c.String(200, "pong")
    })
  
    m := autocert.Manager{
      Prompt:     autocert.AcceptTOS,
      HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
      Cache:      autocert.DirCache("/var/www/.cache"),
    }
  
    log.Fatal(autotls.RunWithManager(r, &m))
  }
  ```

  