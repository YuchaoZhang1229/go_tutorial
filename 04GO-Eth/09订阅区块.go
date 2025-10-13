package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 1. 建立连接, 需要用websocket协议（wss://开头）
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}
	// 2. 创建订阅通道
	headers := make(chan *types.Header)
	// 3. 发起订阅
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("区块头哈希", header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("区块哈希", block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println("区块号", block.Number().Uint64())     // 3477413
			fmt.Println("时间戳", block.Time())                // 1529525947
			fmt.Println("Nonce", block.Nonce())                // 130524141876765836
			fmt.Println("交易数量", len(block.Transactions())) // 7
		}
	}
}
