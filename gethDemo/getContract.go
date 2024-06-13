package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"v1/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()
	conn, err := ethclient.Dial("https://data-seed-prebsc-1-s1.bnbchain.org:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	contract, err := contract.NewContract(common.HexToAddress("0x91e321514434dbe1E5eCA45287B28851f6329dF2"), conn)
	if err != nil {
		log.Fatal(err)
	}
	value, err := contract.GetValue(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	chainid, err := conn.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	privatekey, _ := crypto.HexToECDSA("metamask private")
	auth, err := bind.NewKeyedTransactorWithChainID(privatekey, chainid)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := contract.SetValue(auth, big.NewInt(12345))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("0x%x\n", tx.Hash())
	// 阻塞 等待交易成功
	receipt, err := bind.WaitMined(ctx, conn, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("SET SUCCESS")
	} else {
		fmt.Println("fail")
	}

	// 交易成功 会返回回执 否则没有
	// receipt, err := conn.TransactionReceipt(ctx, tx.Hash())

	value, err = contract.GetValue(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
}
