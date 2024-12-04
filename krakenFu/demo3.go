package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Signature computes the authentication string for the given data, nonce, and endpoint path.
func Signature2(data string, nonce string, endpointPath string, secret string) (string, error) {
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

// Accounts returns key information relating to all your accounts which may either be cash accounts or margin accounts.
// This includes digital asset balances, instrument balances, margin requirements, margin trigger estimates and auxiliary information such as available funds, PnL of open positions and portfolio value.
func Accounts() (map[string]any, error) {
	endpointPath := "/api/v3/accounts"
	endpointURL := EnvironmentURL + "/derivatives" + endpointPath
	signature, err := Signature("", "", endpointPath, PrivateKey)
	fmt.Println(signature)
	if err != nil {
		return nil, fmt.Errorf("failed to compute signature: %s", err)
	}
	req, err := http.NewRequest("GET", endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("APIKey", PublicKey)
	req.Header.Set("Authent", signature)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %s", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err)
	}
	var respBodyMap map[string]any
	if err = json.Unmarshal(respBody, &respBodyMap); err != nil {
		return nil, fmt.Errorf("failed to json unmarshal: %s: %s", err, string(respBody))
	}
	return respBodyMap, nil
}

func main3() {
	accounts, err := Accounts()
	if err != nil {
		log.Fatalf("Error getting accounts: %s", err)
	}
	accountsBytes, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		log.Fatalf("Error JSON encoding accounts: %s", err)
	}
	log.Printf("Accounts: %s", string(accountsBytes))
}
