package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "github.com/go-tutorial/04GO-Eth/store" // for demo
)

func main() {
	// 1. 连接区块链节点
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.GenerateKey()
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// privateKeyHex := hex.EncodeToString(privateKeyBytes)
	// fmt.Println("Private Key:", privateKeyHex)

	// 2. 加载私钥
	privateKey, err := crypto.HexToECDSA("<YOUR_PRIVATE_KEY>")
	if err != nil {
		log.Fatal(err)
	}
	// 3. 推导公钥与地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 4. 设置交易参数 nonce, gasPrice, chainId
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 5. 创建交易认证器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// 6. 部署合约
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("合约部署到", address.Hex())  // 0xE621Ade5a2C0080D8ed406d3504E383eF446ba7B
	fmt.Println("交易哈希", tx.Hash().Hex()) // 0x516f256a7f52d283ce58e4603151ee1a3336468bbc76453f4f4f58525deef472

	_ = instance // 合约实例可以通过该实例来调用合约的方法
}
