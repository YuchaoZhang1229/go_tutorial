package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	contractAddr := "0x943b0324A1B2C5825221D0007951469a492DC8db"
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/jHjyr-bsMbTa4uB-gfigcGghwCvG_iSz")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("a02a2b50874a4f6740500f92b524465c354236641ca49b584bed6704191f72d4")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 估算 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// ------------------------替换代码------------------------
	// 准备交易数据
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)
	// ------------------------替换代码------------------------

	// 创建交易并签名交易
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 300000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	receipt, err := waitForReceipt1(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	if receipt != nil {
		fmt.Println("receipt.Status:", receipt.Status)
	}

	// ------------------------替换代码------------------------
	// 准备查询数据
	itemsSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemsSignature)[:4]

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}
	// ------------------------替换代码------------------------

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// ------------------------替换代码------------------------
	// 解析返回值
	var unpacked [32]byte
	copy(unpacked[:], result)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
	// ------------------------替换代码------------------------
}

func waitForReceipt1(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
