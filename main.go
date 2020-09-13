package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pavelkrolevets/chain_crash/src/client/src/client"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Addr1 string `yaml:"Address1"`
	PK1 string `yaml:"PK1"`
	Addr2 string `yaml:"Address2"`
	PK2 string `yaml:"PK2"`
	Addr3 string `yaml:"Address3"`
	PK3 string `yaml:"PK3"`
	Addr4 string `yaml:"Address4"`
	PK4 string `yaml:"PK4"`
}

var GlobalConfig Config

func (c *Config) getConfig() *Config  {
	confFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(confFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}


func main(){
	GlobalConfig.getConfig()
	log.Println(GlobalConfig)
	// Parse args
	tps:= flag.Int("tps", 1, "tps speed from the node")
	test_size := flag.Int("test_size", 60000, "transaction test size")
	rpc_addr := flag.String("rpc_addr", "http://127.0.0.1:8545", "Node RPC address and open port")
	check_txs :=  flag.Bool("check_txs", false,"Check if all sent transactions are in the blockchain ")
	//threads := flag.Int("threads", 4,"Number of threads to run")
	heavy_io :=  flag.Bool("heavy_io", false,"Add 2048 digit HEX of random data to each tx for heavy io")

	flag.Parse()

	eth_client, err := ethclient.Dial(*rpc_addr)
	if err != nil {
		log.Fatal(err)
	}
	tx_ch_1, tx_ch_2, tx_ch_3, tx_ch_4 := make(chan string, 1), make(chan string, 1), make(chan string, 1), make(chan string, 1)
	go client.SendTransactions(eth_client, GlobalConfig.Addr2, GlobalConfig.PK1, tx_ch_1, *tps, *test_size, *check_txs, *heavy_io)
	go client.SendTransactions(eth_client, GlobalConfig.Addr3, GlobalConfig.PK2, tx_ch_2, *tps, *test_size, *check_txs, *heavy_io)
	go client.SendTransactions(eth_client, GlobalConfig.Addr4, GlobalConfig.PK3, tx_ch_3, *tps, *test_size, *check_txs, *heavy_io)
	go client.SendTransactions(eth_client, GlobalConfig.Addr1, GlobalConfig.PK4, tx_ch_4, *tps, *test_size, *check_txs, *heavy_io)
	result_ch := make(chan string, 1)
	go client.GetTransactionsBlockCount(eth_client, result_ch)

	for {
		select {
		case c1:= <- tx_ch_1:
			log.Println("Received from ch1", c1)
		case c2:= <- tx_ch_2:
			log.Println("Received from ch2", c2)
		case c3:= <- tx_ch_3:
			log.Println("Received from ch3", c3)
		case c4:= <- tx_ch_4:
			log.Println("Received from ch4", c4)
		case c5:= <- result_ch:
			log.Println(c5)
		}
	}
}
