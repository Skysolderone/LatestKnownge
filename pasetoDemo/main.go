package main

import (
	"fmt"
	"log"
	"time"

	"github.com/o1egl/paseto"
)

func main() {
	// 在jwt基础上又套了一层
	symmetricKey := []byte("12345678901234567890123456789012") // Must be 32 bytes
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "test",
		Issuer:     "test_service",
		Jti:        "123",
		Subject:    "test_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}

	jsonToken.Set("data", "0618 paseto Demo")
	v2 := paseto.NewV2()

	encrypt, err := v2.Encrypt(symmetricKey, jsonToken, "0618")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encrypt)
	var newJsonToken paseto.JSONToken
	var newFooter string
	err = v2.Decrypt(encrypt, []byte("12345678901234567890123456789012"), &newJsonToken, &newFooter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newJsonToken)
	fmt.Println(newFooter)
}
