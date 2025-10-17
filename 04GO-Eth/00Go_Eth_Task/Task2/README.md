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