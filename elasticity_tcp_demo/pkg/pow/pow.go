package pow

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

const difficulty = 2

type PoW interface {
	GenerateChallege() (string, string)
	VerifyPow(seed, proof string) bool
}
type PoWImpl struct{}

func (p PoWImpl) GenerateChallege() (string, string) {
	seed := fmt.Sprintf("%d", rand.Int63())
	return seed, fmt.Sprintf("%0*d", difficulty, 0)
}

func (p PoWImpl) VerifyPow(seed, proof string) bool {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(seed+proof)))
	expected := fmt.Sprintf("%0*d", difficulty, 0)
	fmt.Printf("Verifying proof: %s | Expected prefix: %s | Hash: %s\n", proof, expected, hash)

	return hash[:difficulty] == expected
}
