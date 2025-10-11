package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// 💎 核心要点总结
// 连接节点：使用 ethclient.Dial通过RPC URL连接以太坊节点是交互的基础。
// 收据查询方式：
// - BlockReceipts, 通过 区块哈希blockHash 或 区块号blockNumber 查询收据
// - TransactionReceipt, 通过 交易哈希 查询查询收据
// 收据关键信息：交易收据中的 Status、Logs、TxHash、TransactionIndex和 ContractAddress等字段对于分析交易结果至关重要

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	// 1. BlockReceipts, 通过 区块哈希blockHash 或 区块号blockNumber 查询收据
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}

	// 2. TransactionReceipt, 通过 交易哈希 查询查询收据
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)                // 1
	fmt.Println(receipt.Logs)                  // []
	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	fmt.Println(receipt.TransactionIndex)      // 0
	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000

}
