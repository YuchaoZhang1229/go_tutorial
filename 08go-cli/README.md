## 一、项目结构概览
```
// 整体项目目录
myapp/
├── cmd/              # 命令行入口
├── config/           # 配置处理代码
├── internal/         # 私有应用代码
├── pkg/              # 公共库代码
├── configs/          # 配置文件
├── go.mod            # 模块定义
└── main.go           # 程序入口

// 依赖管理
main.go → cmd/ → internal/ → pkg/
                    ↓
                 config/ → configs/
```


### 1. CLI命令逻辑
```
cmd/
├── root.go      # 根命令和全局配置
└── server.go    # 服务器启动命令
```
### 2. 配置处理逻辑
```
config/
├── config.go    # 配置结构体和加载逻辑
└── defaults.go  # 默认配置值
```
###  3. 业务逻辑
```
internal/
├── server/      # HTTP服务器逻辑
├── database/    # 数据访问层
└── logger/      # 日志处理
```
### 4. 配置数据
```
configs/
├── config.yaml          # 主配置文件
├── config.dev.yaml      # 开发环境配置
├── config.prod.yaml     # 生产环境配置
└── config.test.yaml     # 测试环境配置
```
### 5. 可复用库代码
```
pkg/
├── utils/        # 通用工具函数
├── middleware/   # HTTP中间件
└── types/        # 公共类型定义
```

## 二、扩展性
当需要添加新功能时：
- 新命令 → 添加到 cmd/
- 新配置 → 扩展到 config/
- 新业务逻辑 → 添加到 internal/
- 新公共组件 → 添加到 pkg/

## 三、使用方式
### 1. 查看帮助
```bash
go run .\main.go help
```

### 2. 启动服务器（使用默认配置）
```bash
go run .\main.go server
```

### 3. 指定配置文件
```bash
go run .\main.go server --config ./configs/config.yaml
```

### 4. 使用命令行参数
```bash
go run .\main.go server --port 9090 --log-level debug
```

### 5. 使用环境变量
```bash
# 环境变量会自动映射（MYAPP_ 前缀）
export MYAPP_LOG_LEVEL=debug
export MYAPP_DATABASE_HOST=localhost
go run .\main.go
```

## 四、主要特性
1. **Cobra**: 处理命令行参数和命令结构 
2. **Viper**: 统一管理配置文件、环境变量、命令行参数
3. **配置优先级**: 命令行参数 > 环境变量 > 配置文件 > 默认值 
4. **模块化设计**: 清晰的目录结构和职责分离 
5. **生产就绪**: 包含日志、数据库、优雅关闭等特性