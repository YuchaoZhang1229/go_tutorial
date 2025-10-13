**下载这两个文件 IERC20.sol, IERC20Metadata.sol:**
- https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol
- https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/IERC20Metadata.sol


**使用`solcjs`工具单独处理，生成调用合约时所需的 ABI 的 JSON 文件，并使用 `abigen` 工具根据 ABI 的 JSON 文件生成 go 代码。**
```go
// 将 npm config get prefix 结果添加到PATH环境变量中
npm install -g solc
solcjs --version

// 将 go env GOPATH 结果添加到PATH环境变量中
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
abigen --version

solcjs --abi IERC20Metadata.sol
abigen --abi=IERC20Metadata_sol_IERC20Metadata.abi --pkg=erc20 --out=erc20.go
```