package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"github.com/astra-x/go-ethereum/common"
	"github.com/astra-x/go-ethereum/crypto"
	"github.com/astra-x/go-ethereum/ethclient"
	//"time"
	"math/big"
	"github.com/astra-x/go-ethereum/core/types"
	"time"
	"flag"
)


func getTransactionCount(client *ethclient.Client, ch chan<-uint)  {
	ticker:= time.NewTicker(time.Second)
	for {
		blockhash, err:= client.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
		proc_txs, err:= client.TransactionCount(context.Background(),blockhash.Hash())
		if err != nil {
			log.Fatal(err)
		}
		select {
		case <-ticker.C:
		}
		ch <- proc_txs
	}
}

func main() {
	tps:= flag.Int("tps", 100, "tps speed from the node")
	test_size := flag.Int("test_size", 60000, "transaction test size")
	rpc_addr := flag.String("rpc_addr", "http://127.0.0.1:8545", "Node RPC address and open port")
	flag.Parse()

	tx_block := make(chan uint, 1)

	client1, err := ethclient.Dial(*rpc_addr)
	if err != nil {
		log.Fatal(err)
	}


	go getTransactionCount(client1, tx_block)



	for {
		select {
		case out := <-tx_block:
			fmt.Println("Transactions in the last block:  ",out)
		}

	}
}
