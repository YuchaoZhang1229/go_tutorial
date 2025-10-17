### 任务 2：合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。

具体任务

1. 编写智能合约
    - 使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
    - 编译智能合约，生成 ABI 和字节码文件。
2. 使用 abigen 生成 Go 绑定代码
    - 安装 abigen 工具。
    - 使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
3. 使用生成的 Go 绑定代码与合约交互
    - 编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
    - 调用合约的方法，例如增加计数器的值。
    - 输出调用结果。


### 安装 sloc和 abigen 工具
```go
// 将 npm config get prefix 结果添加到PATH环境变量中
npm install -g solc
solcjs --version

// 将 go env GOPATH 结果添加到PATH环境变量中
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
abigen --version
```

### 使用 abigen生成 Go 绑定代码
```bash
# 译合约代码，会在当目录下生成一个编译好的二进制字节码文件 store_sol_Store.bin
solcjs --bin Counter.sol

# 生成合约 abi 文件，会在当目录下生成 store_sol_Store.abi 文件
solcjs --abi Counter.sol

# 使用 abigen 工具根据这两个生成 bin 文件和 abi 文件，生成 go 代码
abigen --bin=Counter_sol_Counter.bin --abi=Counter_sol_Counter.abi --pkg=counter --out=counter.go
```