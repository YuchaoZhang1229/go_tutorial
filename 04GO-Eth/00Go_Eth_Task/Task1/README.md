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
