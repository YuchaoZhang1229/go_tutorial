package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

// GetHistoricalLogs 通过过滤器（FilterLogs）获取指定区块范围内的历史事件日志
func GetHistoricalLogs(client *ethclient.Client, contractAddress common.Address) {
	// 构建过滤查询
	query := ethereum.FilterQuery{
		// BlockHash查询某个特定区块哈希

		// FromBlock范围查询, 单独的话就是从指定的区块查询到最新区块
		FromBlock: big.NewInt(9418563),
		//ToBlock:   big.NewInt(9418520),
		Addresses: []common.Address{contractAddress},

		// 可以添加Topics来过滤特定事件
		// Topics: [][]common.Hash{
		//  {},
		//  {},
		// },
	}

	// 获取日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("在区块 %d 到 最新 范围内找到 %d 条日志\n",
		big.NewInt(6920583), len(logs))

	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		log.Fatal(err)
	}
	// 解析日志
	for _, vLog := range logs {
		fmt.Println("日志索引:", vLog.Index)
		fmt.Println("区块哈希:", vLog.BlockHash.Hex())
		fmt.Println("区块号:", vLog.BlockNumber)
		fmt.Println("交易哈希:", vLog.TxHash.Hex())

		// 进一步解析日志...
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("key:", common.Bytes2Hex(event.Key[:]))
		fmt.Println("value", common.Bytes2Hex(event.Value[:]))
		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex())
		}

		fmt.Println("topics[0]=", topics[0])
		if len(topics) > 1 {
			fmt.Println("indexed topics (key):", topics[1:])
		}
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("signature topics=", hash.Hex(), "与topics[0]相同") // 应该与topics[0]相同
}

// ListenEvents 通过订阅（SubscribeFilterLogs）实时监听事件, 必须要要用websocket wss
func ListenEvents(client *ethclient.Client, contractAddress common.Address) {
	// 构建订阅查询
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 创建日志通道
	logs := make(chan types.Log)

	// 订阅事件
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("订阅失败: %v", err)
	}

	fmt.Println("开始监听事件...")

	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("订阅错误: %v", err)
		case vLog := <-logs:
			fmt.Println("日志索引:", vLog.Index)
			fmt.Println("区块哈希:", vLog.BlockHash.Hex())
			fmt.Println("区块号:", vLog.BlockNumber)
			fmt.Println("交易哈希:", vLog.TxHash.Hex())

			// 进一步解析日志...
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("key:", common.Bytes2Hex(event.Key[:]))
			fmt.Println("value", common.Bytes2Hex(event.Value[:]))
			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}

			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("indexed topics (key):", topics[1:])
			}
		}
	}
}

func main() {
	client, err := ethclient.Dial("wss://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xbc210db44E91520D35350f05f216734638D278Eb")
	//GetHistoricalLogs(client, contractAddress)
	ListenEvents(client, contractAddress)
}
