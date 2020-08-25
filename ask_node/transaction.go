package main

import (
	"context"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"

	"time"
	"flag"
	"fmt"
)


func getTransactionCount(client *ethclient.Client)  {

	ticker:= time.NewTicker(time.Second)

	for {

		blockhash, err:= client.BlockByNumber(context.Background(),nil)

		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}

		proc_txs, err:= client.TransactionCount(context.Background(),blockhash.Hash())

		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		select {
		case <-ticker.C:
			//ch <- proc_txs
			fmt.Println("Transactions in the block #", blockhash.NumberU64(), " :", proc_txs, "\n",
				"Block time :", blockhash.Time(),
					"Validator :", blockhash.Header().Coinbase.Hex())
			}

	}
}

func main() {
	rpc_addr := flag.String("rpc_addr", "http://127.0.0.1:8545", "Node RPC address and open port")
	flag.Parse()

	//tx_block := make(chan uint, 1)

	client1, err := ethclient.Dial(*rpc_addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		getTransactionCount(client1)
	}


	//for {
	//	select {
	//	case out := <-tx_block:
	//		fmt.Println("Transactions in the last block:  ",out)
	//	}
	//
	//}
}
