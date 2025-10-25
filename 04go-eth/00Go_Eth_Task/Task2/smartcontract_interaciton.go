package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	counter "github.com/go-tutorial/04GO-Eth/00Go_Eth_Task/Task2/counter"
)

func deployContract(client *ethclient.Client, privateKey string) (common.Address, *counter.Counter) {
	// 加载私钥
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 获取公钥
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	// 获取账户地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("账户地址:", fromAddress.Hex())
	// 设置交易参数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 部署合约
	address, tx, Counter, err := counter.DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("部署合约合约地址:", address.Hex())
	fmt.Println("部署合约交易哈希:", tx.Hash().Hex())

	// 等待部署成功
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("部署合约交易成功:", receipt.Status)

	return address, Counter
}

// 调用合约
func callContract(client *ethclient.Client, address common.Address, privateKey string) {
	// 创建合约实例
	Counter, err := counter.NewCounter(address, client)
	if err != nil {
		log.Fatal(err)
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 创建交易
	opt, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := Counter.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Increment 交易哈希:", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("等待计数器+1成功:", receipt.Status)

	// 查询计数
	callOpts := &bind.CallOpts{Context: context.Background()}
	count, err := Counter.GetCount(callOpts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("计数:", count.Int64())
}

func main() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	privateKey := "<YOUR_PRIVATE_KEY>"
	//address, _ := deployContract(client, privateKey)

	callContract(client, common.HexToAddress("0xfA485B4FA50317822987666cA2f1ee2F222cdf48"), privateKey)
}
