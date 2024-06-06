package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	conn, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	// block header
	header, err := conn.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())
	// seacrh block
	blockNum := big.NewInt(0)
	block, err := conn.BlockByNumber(ctx, blockNum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block)

	// 返回block的交易数量
	count, err := conn.TransactionCount(ctx, block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
