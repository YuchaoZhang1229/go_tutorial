package main

import (
	"context" // 用于控制调用上下文（如超时、取消）
	"fmt"
	"log" // 处理大整数（以太坊中很多数值都是大整数）
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient" // go-ethereum库的客户端，用于连接以太坊节点并与之交互
)

// 01查询区块
// 1. 连接以太坊节点:
// 通过 Alchemy 提供的 Sepolia 测试网 API 端点建立与以太坊节点的连接。
// 2. 查询区块信息:
// 获取指定区块号（5671744）的：
// ● 区块头信息（Header）
// ● 完整区块信息（Block）
// ● 区块内交易数量
// 3. 打印关键数据:
// 输出区块号、时间戳、难度值、区块哈希和交易数量等关键信息。

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)

	// 1. 查询区块头信息（Header）
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1. client.HeaderByNumber: 查询区块头信息（Header）")
	fmt.Println("区块号: ", header.Number.Uint64())     // 5671744
	fmt.Println("时间戳: ", header.Time)                // 1712798400
	fmt.Println("难度值: ", header.Difficulty.Uint64()) // 0
	fmt.Println("区块哈希: ", header.Hash().Hex())       // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(" ")
	// 2. 查询完整区块信息（Block）
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2. client.BlockByNumber: 查询完整区块信息（Block）")
	fmt.Println("区块号: ", block.Number().Uint64())     // 5671744
	fmt.Println("时间戳: ", block.Time())                // 1712798400
	fmt.Println("难度值: ", block.Difficulty().Uint64()) // 0
	fmt.Println("区块哈希: ", block.Hash().Hex())         // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println("交易数量: ", len(block.Transactions()))  // 70
	fmt.Println(" ")

	// 3. 查询区块内交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("3. client.TransactionCount: 查询区块内交易数量")
	fmt.Println("交易数量: ", count) // 70
	fmt.Println(" ")

}
