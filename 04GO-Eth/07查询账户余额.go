package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xE621Ade5a2C0080D8ed406d3504E383eF446ba7B")

	// 查询当前账户余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BalanceAt:", balance) // 169714560621028552

	// 查询指定区块高度下的账户余额
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("BalanceAt with blockNumber:", balanceAt) // 0

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue:", ethValue) //  0.169714560621028552

	PendingBalanceAt, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("PendingBalanceAt:", PendingBalanceAt) // 169714560621028552
}
