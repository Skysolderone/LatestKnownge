package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	AccessKey = "oTkVVDH9kfmwv1Ilj3N4ev2ocP4uwQzzoaUeaivgI3pqgpN7bL9YjNG9"
	SecretKey = "O8eJv6Crz87f2mnnaJCD/jlFj9KlQJ0Jy41DEP6boju6NX6kCBoj5bw9z66/5HYbUWm+TQ9M6wS6aO4mN08z4Z+6"
)

func Signature(data string, nonce string, endpointPath string, secret string) (string, error) {
	message := data + nonce + endpointPath
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(message))
	messageHashed := sha256Hash.Sum(nil)
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to decode secret: %s", err)
	}
	hmacHash := hmac.New(sha512.New, decodedSecret)
	hmacHash.Write(messageHashed)
	signature := hmacHash.Sum(nil)
	b64Signature := base64.StdEncoding.EncodeToString(signature)
	return b64Signature, nil
}

func main() {
	url := "https://futures.kraken.com/derivatives/api/v3/sendorder"
	method := "POST"

	payload := strings.NewReader(`{
		"ProcessBefore": "2023-11-08 19:56:35.441899+00:00",
		"orderType": "mkr",
		"symbol": "string",
		"side": "buy",
		"size": 0,
		"limitPrice": 0,
		"stopPrice": 0,
		"cliOrdId": "string",
		"triggerSignal": "mark",
		"reduceOnly": true,
		"trailingStopMaxDeviation": 0,
		"trailingStopDeviationUnit": "PERCENT",
		"limitPriceOffsetValue": 0,
		"limitPriceOffsetUnit": "QUOTE_CURRENCY"
	  }`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, err := Signature("", "", "/api/v3/sendorder", SecretKey)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("APIKey", AccessKey)
	req.Header.Add("Authent", signature)

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
