package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/go-tutorial/04GO-Eth/store"
)

func main() {
	// 1. 创建 ethclient 实例
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	// 2. 创建合约实例
	contractAddr := "0x943b0324A1B2C5825221D0007951469a492DC8db"
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	// 3. 根据hex创建私钥实例
	privateKey, err := crypto.HexToECDSA("<YOUR_PRIVATE_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	// 4. 调用合约方法
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))
	// a. 用于发送交易（如调用合约函数修改状态）
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())
	// b. 查询合约
	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)

	version, err := storeContract.Version(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("version:", version)
}
