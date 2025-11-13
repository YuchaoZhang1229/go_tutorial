## Go Modules（Go 1.11+）
- 官方推出的依赖管理方案
- 不再需要 GOPATH
- 成为现在的标准方案

### 1. 基本命令
```bash
# 初始化模块
go mod init github.com/<module-name>

# 添加依赖
go get [package]
go get [package]@[version]

# 整理依赖
go mod tidy

# 验证依赖
go mod verify
```

```bash
# 查看依赖图
go mod graph

# 清理缓存
go mod clean
```

### go.mod文件结构
```
// 定义模块路径
module github.com/<module-name>

// 指定项目使用的 Go 版本
go 1.19

// // replace 指令用于替换依赖包
replace github.com/old/pkg => github.com/new/pkg v1.0.0

// require 指令声明项目所需的直接依赖
// 推荐为每个依赖指定明确的版本号，避免使用不明确的版本
require (
    github.com/gin-gonic/gin v1.8.1
    github.com/go-sql-driver/mysql v1.6.0
)   

// exclude 指令用于排除特定的依赖版本
exclude (
    github.com/problematic/pkg v1.2.3
    github.com/another/bad v0.5.0
)
```