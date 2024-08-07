package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip39"
)

// github.com/tyler-smith/go-bip39库来处理助记词和种子的转换
func main() {
	// 生成随机熵
	encrypt, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatal(err)
	}
	// 生成助记词
	mnemonic, err := bip39.NewMnemonic(encrypt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mnemonic :%s\n", mnemonic)
	// 从助记词生成种子Seed
	seed := bip39.NewSeed(mnemonic, "0807") // phrase
	fmt.Printf("seed :%x\n", seed)
	// 从种子生成主密钥
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Master Key: %v\n", masterKey)
	// 派生子密钥并生成比特币地址
	for i := 0; i < 5; i++ {
		childKey, err := masterKey.Derive(uint32(i))
		if err != nil {
			log.Fatal(err)
		}
		// 转换位公钥
		pubKey, err := childKey.Neuter()
		if err != nil {
			log.Fatalf("Failed to neuter child key: %v", err)
		}
		// 生成公钥地址
		address, err := pubKey.Address(&chaincfg.MainNetParams)
		if err != nil {
			log.Fatalf("Failed to generate address: %v", err)
		}

		fmt.Printf("Address %d: %s\n", i, address.EncodeAddress())

	}
}
