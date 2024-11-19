package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main2() {
	url := "https://futures.kraken.com/derivatives/api/v3/accounts"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	auth, err := generateHMAC(Sec, "/api/v3/accounts", nonce, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKey", Api)
	req.Header.Add("Authent", auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func generateHMAC(secret, urlPath, nonce, data string) (string, error) {
	// Decode the base64-encoded secret
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// Create the SHA256 hash of the nonce and data
	hash256 := sha256.New()
	hash256.Write([]byte(nonce + data))
	hashedData := hash256.Sum(nil)

	// Create the HMAC using SHA512
	var mac hash.Hash = hmac.New(sha512.New, decodedSecret)
	mac.Write([]byte(urlPath))
	mac.Write(hashedData)
	hmacResult := mac.Sum(nil)

	// Encode the HMAC result to base64
	encodedHMAC := base64.StdEncoding.EncodeToString(hmacResult)

	return encodedHMAC, nil
}

func SignParam(urlPath, secret, payload string) (signStr string, err error) {
	sha := sha256.New()
	sha.Write([]byte(payload))
	shasum := sha.Sum(nil)

	s, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return
	}
	mac := hmac.New(sha512.New, s)
	mac.Write(append([]byte(urlPath), shasum...))
	mac.Write(shasum)
	macsum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macsum), nil
}
