package main

import (
	"log"
	"github.com/astra-x/go-ethereum/ethclient"
	"context"
	"fmt"
	"math/big"
	"github.com/astra-x/go-ethereum/common"
)

func getBlock(hash common.Hash) {

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	block_num := big.NewInt(1210)

	block, _ := client.BlockByNumber(context.Background(), block_num)
	header, _ := client.HeaderByNumber(context.Background(), block_num)
	transaction, _, _:= client.TransactionByHash(context.Background(), hash)
	fmt.Println(block, header, transaction)
}


func main() {
	hash:= common.HexToHash("0xd1529c6f7edd76f9ab4b02c9f7bf3d337dd66bc0f4a386fe6bdbbe78044f8290")
	getBlock(hash)
}
