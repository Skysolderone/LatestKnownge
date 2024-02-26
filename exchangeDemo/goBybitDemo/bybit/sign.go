package bybit

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SignParam(timestamp, recvWindow, paramStr, apikey, apiSecret string) (signStr string) {
	payload := timestamp + apikey + recvWindow + paramStr
	signStr, _ = HmacSHA256Base64Sign(apiSecret, payload)
	return
}

func HmacSHA256Base64Sign(secret, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	signByte := mac.Sum(nil)
	return hex.EncodeToString(signByte), nil
}
