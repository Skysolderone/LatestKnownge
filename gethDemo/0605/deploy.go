package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const key = ""

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/")
	if err != nil {
		log.Fatalf("failed to connect to the ethereumclient :%v", err)
	}
	chainId, err := conn.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainId :%v", err)
	}
	// _, err = bind.NewTransactorWithChainID(strings.NewReader(key), "password", chainId)
	// if err != nil {
	// 	log.Fatalf("bing chain :%v", err)
	// }
	fmt.Println(chainId)
	// address, tx, instance, err := DeployStorage(auth, conn)
	// address, tx, _, err := DeployStorage(auth, conn)
	// if err != nil {
	// 	log.Fatalf("deploy contract %v", err)
	// }
	// fmt.Printf("Contract deploy pending :0x%v", address)
	// fmt.Printf("Transaction be mined :0x%x\n\n", tx.Hash())
	// time.Sleep(250 * time.Millisecond)
	// 合约调用
	// instance.Store()
	// instance.Retrieve()
	// name, err :=
	// if err != nil {
	// 	log.Fatalf("Failed to retrived pending name :%v", err)
	// }
	// fmt.Println("Pending name :", name)
}
