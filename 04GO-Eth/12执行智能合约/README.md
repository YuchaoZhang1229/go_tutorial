## 一、使用 abigen 工具生成的代码




## 二、仅使用 ethclient 包调用
### 1. 使用 abi 文件
#### 🔎 **代码模块详解**
##### 1. 连接区块链网络
这部分代码负责与以太坊网络建立一个连接通道。
```go
client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<YOUR_API_KEY>")
if err != nil {
    log.Fatal(err)
}
```
- ethclient.Dial：函数用于建立一个与以太坊节点的远程过程调用（RPC）连接。
- 网络地址：代码中使用的 https://...是连接到 Sepolia 测试网的节点服务。测试网用于开发和测试，不使用真实的有价货币

##### 2. 地址推导
这部分代码处理交易发送者的身份验证。
go
```go
privateKey, err := crypto.HexToECDSA("你的私钥")
publicKey := privateKey.Public()
publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
```

- **私钥**：是控制账户资产和权限的最高凭证，**绝对不可以泄露**。在生产环境中，应使用安全的密钥管理系统（如加密文件或硬件钱包）替代代码中的明文私钥。
- **地址推导**：以太坊地址是从私钥对应的公钥经过哈希计算得出的，可以公开分享，用于接收资产

##### 3. 构造交易数据
```go
nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
gasPrice, err := client.SuggestGasPrice(context.Background())
```
- **Nonce**：是一个账户发出的交易序列号，用于确保交易顺序且防止重放。每发送一笔新交易，Nonce 值加 1
- **Gas Price 与 Gas Limit**：
  - **Gas** 是衡量在以太坊上执行操作所需计算工作量的单位。
  - **Gas Price** 是你愿意为每个 Gas 单位支付的价格。价格越高，交易被矿工优先打包的可能性越大
  - **Gas Limit** 是你愿意为这笔交易消耗的最大 Gas 数量（代码中硬编码为 300000）。设置过低可能导致交易失败（但已消耗的Gas不退还）


##### 4. 编码交易数据 (与合约交互的核心)
这是与智能合约交互最关键的一步，即如何告诉合约“你想调用哪个函数”以及“传递什么参数”。
```go
contractABI, err := abi.JSON(strings.NewReader(`[...]`)) // 1. 加载ABI
input, err := contractABI.Pack("setItem", key, value) // 2. 打包数据
```
**ABI 打包**：abi.Pack函数根据 ABI 定义，将函数名 ```setItem``` 和参数 ```key```, ```value``` 编码成以太坊虚拟机可以理解的二进制数据（input）。这笔数据将作为交易 Data字段的内容


##### 5. 创建、签名与发送交易 (交易)
将前述所有部分组合成一笔完整的交易，并用私钥签名后广播到网络。
```go
tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 300000, gasPrice, input)
signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
err = client.SendTransaction(context.Background(), signedTx)
```
- **交易创建**：```types.NewTransaction```创建一笔交易。其中 big.NewInt(0)表示不发送以太币，交易的目标地址是合约地址
- **交易签名**：使用私钥对交易进行签名，以证明你有权从该账户发起此交易
- **waitForReceipt函数**：这是一个自定义的循环等待函数，用于等待交易被网络确认并获取收据。交易收据包含了交易的最终执行结果（如状态、Gas 实际消耗量等）

##### 6. 查询合约状态（调用）
在交易确认后，通过调用的方式查询合约状态，验证数据是否已正确写入。 
```go
callInput, err := contractABI.Pack("items", key) // 1. 打包查询请求
callMsg := ethereum.CallMsg{To: &to, Data: callInput} // 2. 构造调用消息
result, err := client.CallContract(context.Background(), callMsg, nil) // 3. 执行调用
var unpacked [32]byte
contractABI.UnpackIntoInterface(&unpacked, "items", result) // 4. 解析结果
```
- **调用**：client.CallContract在本地节点执行合约代码，不会产生交易，也不消耗 Gas 。它直接返回函数执行的结果
#### 💡 关键知识点与交互模式对比
| **特性**             | **交易**                    | **调用**        |
|----------------|-----------------------|-----------|
| **操作类型**           | 写操作（修改状态）             | 读操作（查询状态） |
| **执行位置**           | 在全网节点共识后执行                   | 在连接的本地节点执行      |
| **Gas 费用**         | 需要支付                  | 免费        |
| **链上效果**           | 改变区块链状态               | 不改变状态     |
| **返回值**            | 不直接返回，需通过事件日志或后续查询获取  | 直接返回结果    |
| **代码中的方法**         | SendTransaction | CallContract |

### 2. 不使用 abi 文件

