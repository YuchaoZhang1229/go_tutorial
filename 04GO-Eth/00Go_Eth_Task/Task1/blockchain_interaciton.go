package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	// 1. 查询区块
	// https://sepolia.etherscan.io/tx/0xb829637795488c81c60156d7dcc81e9d8dc84f2534fdf2e786d79dd192d29bd8
	block, err := client.BlockByNumber(context.Background(), big.NewInt(9418577))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("区块高度:", block.Number().Uint64())
	fmt.Println("区块哈希:", block.Hash().Hex())
	fmt.Println("时间戳:", block.Time())
	fmt.Println("交易数量:", len(block.Transactions()))

	// 2. 发送交易
	privateKey, err := crypto.HexToECDSA("a02a2b50874a4f6740500f92b524465c354236641ca49b584bed6704191f72d4")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	value := big.NewInt(1000000000000000) // 0.001 ETH in wei
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	toAddress := common.HexToAddress("0x0D9c3E1F8Dc9fc4E46BfE724c565803CeE3b6E6B")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("交易哈希", signedTx.Hash().Hex())

	// 等待交易被挖出（这里等待5分钟超时）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, client, signedTx)
	if err != nil {
		log.Fatalf("等待交易确认时出错: %v", err)
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Printf("交易成功确认！区块号: %d\n", receipt.BlockNumber.Uint64())
	} else {
		fmt.Printf("交易执行失败！\n")
	}
}
