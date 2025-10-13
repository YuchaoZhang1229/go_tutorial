package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "github.com/go-tutorial/04GO-Eth/store" // for demo
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xE621Ade5a2C0080D8ed406d3504E383eF446ba7B")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
