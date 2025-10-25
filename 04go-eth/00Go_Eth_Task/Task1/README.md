

### 任务 1：区块链读写 任务目标

使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。

具体任务

1. 环境搭建
    - 安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
    - 注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
2. 查询区块
    - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
    - 实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
    - 输出查询结果到控制台。
3. 发送交易
    - 准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
    - 编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
    - 构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
    - 对交易进行签名，并将签名后的交易发送到网络。
    - 输出交易的哈希值。

### ⚠️ 常见问题与处理策略
在交易的生命周期中，可能会遇到以下几种状态，了解它们有助于你更好地排查问题
1. **Pending**: 交易在内存池中等待打包。如果长时间未打包，可能是 GasPrice 过低，可尝试用**相同 Nonce** 和**更高 GasPrice** 重新发送一笔交易来替换它
2. **Failed**: 交易已打包但执行失败（例如合约调用 revert）, **Gas 费会被扣除**, 需要通过交易收据解析失败原因
3. **Dropped**: 交易因 GasPrice 过低等原因被从内存池中移除。需要**提高 GasPrice 后重新发送**。

### 💡 进阶提示：EIP-1559 交易
除了上面使用的传统交易类型（Legacy Transaction），目前更推荐使用 EIP-1559 类型的交易，它提供了更好的费用预测体验 。它使用 GasFeeCap（你愿意支付的最高总费用）和 GasTipCap（给矿工的小费）来代替单一的 GasPrice。

```go
// 构建 EIP-1559 交易（DynamicFeeTx）
gasTipCap, _ := client.SuggestGasTipCap(context.Background())
gasFeeCap, _ := client.SuggestGasPrice(context.Background()) // 简化处理，实际可计算基础费+小费

tx := types.NewTx(&types.DynamicFeeTx{
    ChainID:   chainID,
    Nonce:     nonce,
    GasTipCap: gasTipCap,
    GasFeeCap: gasFeeCap,
    Gas:       gasLimit,
    To:        &toAddress,
    Value:     value,
    Data:      nil,
})
// ...（后续的签名和发送步骤与传统交易相同）
```
