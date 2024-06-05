package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("dial failed %v", err)
	}
	defer conn.Close()
	chainid, _ := conn.ChainID(ctx)
	if common.IsHexAddress("0x6E98D09D239a0F0AC53081C92e338eB78a04215C") {
		fmt.Println("IS TURE")
	}
	address := common.HexToAddress("0x6E98D09D239a0F0AC53081C92e338eB78a04215C")
	balacen, _ := conn.BalanceAt(ctx, address, nil) //
	fmt.Println(chainid)
	fmt.Println(balacen)
	// wei to ether
	fbalance := new(big.Float)
	fbalance.SetString(balacen.String())
	ethvalue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethvalue)
	pendingBalace, _ := conn.PendingBalanceAt(ctx, address)
	fmt.Println(pendingBalace)
}
