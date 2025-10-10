package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 初始化以太坊客户端连接
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	// 2. 获取网络链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 3. 通过区块号查询完整的区块信息
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 4. 遍历区块中的交易并提取关键信息
	for _, tx := range block.Transactions() {
		// 交易哈希（Hash）：交易的唯一标识符。
		// 转账金额（Value）：交易转移的以太坊原生代币（ETH）数量。
		// Gas相关参数（Gas, GasPrice）：Gas是计算工作的单位，GasPrice是单位Gas的价格，两者共同决定了交易手续费。
		// Nonce：由发送账户发出的交易序列号，用于防止双花和确保交易顺序。
		// 数据（Data）：通常用于调用智能合约时的输入参数。
		// 接收方地址（To）：交易的目标地址。
		fmt.Println("----------------------------交易信息----------------------------")
		fmt.Println("交易哈希: ", tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println("转账金额: ", tx.Value().String())    // 100000000000000000
		fmt.Println("Gas: ", tx.Gas())                    // 21000
		fmt.Println("GasPrice: ", tx.GasPrice().Uint64()) // 100000000000
		fmt.Println("Nonce: ", tx.Nonce())                // 245132
		fmt.Println("Data: ", tx.Data())                  // []
		fmt.Println("To: ", tx.To().Hex())                // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// 5. 从已签名的交易（Signed Transaction）中恢复出生成该签名的以太坊账户地址（即交易发送方）
		// func Sender(signer Signer, tx *Transaction) (common.Address, error)
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}

		// 6. 查询交易收据
		// 接收交易哈希
		// Status（0表示失败，1表示成功）和Logs（智能合约执行过程中产生的事件日志）
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("receipt.Status: ", receipt.Status) // 1
		fmt.Println("receipt.Logs: ", receipt.Logs)     // []
		fmt.Println("-------------------------------------------------------------")
		break
	}

	// 8. 根据区块哈希和交易索引位置查询特定交易
	fmt.Println("区块哈希: ", block.Hash().Hex())
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("交易哈希: ", tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	// 9. 通过交易哈希来查询交易
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("isPending: ", isPending)      // isPending参数可以指示该交易是否还在等待被纳入区块的状态（待处理）
	fmt.Println("交易哈希: ", tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false
}
