package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 24

type Block struct {
	PrevBlockHash string
	Data          string
	Nonce         int
}

func main() {
	block := Block{
		PrevBlockHash: "0000000000000000000000000000000000000000000000000000000000000000",
		Data:          "Hello, Blockchain!",
		Nonce:         0,
	}

	// 设置目标值
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	// 开始挖矿
	for block.Nonce < math.MaxInt64 {
		data := block.prepareData()
		hash := sha256.Sum256([]byte(data))

		hashInt := new(big.Int)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(target) == -1 {
			fmt.Printf("Nonce found: %d\n", block.Nonce)
			fmt.Printf("Hash: %s\n", hex.EncodeToString(hash[:]))
			break
		} else {
			block.Nonce++
		}
	}
}

func (b *Block) prepareData() string {
	return fmt.Sprintf("%s%s%d", b.PrevBlockHash, b.Data, b.Nonce)
}
