package main

import (
"log"
"github.com/astra-x/go-ethereum/ethclient"

	"fmt"
	"context"
	"time"
)

func check_pool() {

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	pending,_ := client.PendingTransactionCount(context.Background())
	fmt.Printf("\rPending %d", pending)

	//txpool := core.TxPool{}
	//pending, queued := txpool.Stats()
	//fmt.Println("Pending :",pending,"Queued :", queued)
}


func main() {

	for {
		time.Sleep(1 * time.Second)
		check_pool()

	}
}

