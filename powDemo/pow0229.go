package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Blocktwo struct {
	Nonce int
	Data  string
	Hash  string
}

// copilot
func main() {
	data := "Hello world"
	diffculty := 4
	for nonce := 0; ; nonce++ {
		block := Blocktwo{Nonce: nonce, Data: data}
		hash := block.calculateHash()
		fmt.Printf("\r%x", hash)
		if strings.HasPrefix(hash, strings.Repeat("0", diffculty)) {
			block.Hash = hash
			fmt.Printf("\n\n%s", block)
			break
		}
	}
}

func (b *Blocktwo) calculateHash() string {
	record := strconv.Itoa(b.Nonce) + b.Data
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

func (b *Blocktwo) String() string {
	return fmt.Sprintf("Data: %s\nHash: %s\nNonce: %d", b.Data, b.Hash, b.Nonce)
}
