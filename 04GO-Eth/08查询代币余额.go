package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/go-tutorial/04GO-Eth/erc20" // for demo
	"log"
	"math"
	"math/big"
)

// 下载这两个文件
// https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/IERC20.sol
// https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/extensions/IERC20Metadata.sol

// 将 npm config get prefix 结果添加到PATH环境变量中
// npm install -g solc
// solcjs --version

// 将 go env GOPATH 结果添加到PATH环境变量中
// go install github.com/ethereum/go-ethereum/cmd/abigen@latest
// abigen --version

// solcjs --abi IERC20Metadata.sol
// abigen --abi=IERC20Metadata_sol_IERC20Metadata.abi --pkg=token --out=erc20.go

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/jHjyr-bsMbTa4uB-gfigcGghwCvG_iSz")
	if err != nil {
		log.Fatal(err)
	}
	// 填写代币地址
	tokenAddress := common.HexToAddress("0xF2758163DF55c7BC9411bD5Ff5F79ed2C4C12d7C")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	// 填写钱包地址
	address := common.HexToAddress("0xE621Ade5a2C0080D8ed406d3504E383eF446ba7B")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
