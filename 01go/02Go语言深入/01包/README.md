## Go 语言包（Package）知识点详解

### 1. 包声明和导入
#### 包声明
```go
// 每个 Go 文件必须以包声明开始
package main

// 或者
package utils
```

#### 包导入
```go
import "fmt"
import "math/rand"

// 分组导入
import (
    "fmt"
    "math"
    "strings"
)

// 带别名的导入
import (
    f "fmt"
    m "math"
)

// 点导入（不推荐）
import . "fmt"  // 可以直接使用 Println 而不需要 fmt.Println

// 空白标识符导入（只执行包的 init 函数）
import _ "database/sql"
```

### 2. init 函数
#### init 函数特性
- 每个包可以包含多个 init 函数
- init 函数在程序开始时自动执行
- 执行顺序：导入的包 → 当前包的常量/变量 → init 函数
```go
package config

import "fmt"

var config map[string]string

func init() {
    config = make(map[string]string)
    config["host"] = "localhost"
    config["port"] = "8080"
    fmt.Println("配置初始化完成")
}

func init() {
    // 可以有多个 init 函数
    fmt.Println("第二个初始化函数")
}
```

### 3. 包初始化顺序
初始化流程
1. 导入的包按依赖顺序初始化
2. 包级别的常量和变量初始化
3. 执行 init 函数
4. main 包初始化
5. 执行 main 函数

```go
package main

import "fmt"

var globalVar = func() int {
    fmt.Println("全局变量初始化")
    return 100
}()

func init() {
    fmt.Println("init 函数执行")
}

func main() {
    fmt.Println("main 函数执行")
}
```

### 4. 内部包（internal）
internal 包的特殊性
- internal 目录下的包只能被父目录下的包导入
- 提供了一种包级别的访问控制机制
```
myproject/
├── main.go
└── internal/
    └── auth/
        └── auth.go  // 只能被 myproject 下的包导入
```

### 5. 常用标准包

```go
import (
    "fmt"       // 格式化 I/O
    "os"        // 操作系统功能
    "io"        // I/O 原语
    "bufio"     // 缓冲 I/O
    "strings"   // 字符串操作
    "strconv"   // 字符串转换
    "math"      // 数学函数
    "time"      // 时间操作
    "encoding/json" // JSON 编码解码
    "net/http"  // HTTP 客户端和服务器
)
```
