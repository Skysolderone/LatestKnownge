package config

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))
	return hex.EncodeToString(h.Sum(nil))
}

func getParameters(dataStr string, payload string, urlEncode bool) map[string]interface{} {
	var apiMap map[string]interface{}
	var payloadMap map[string]interface{}
	timestemp := time.Now().UnixNano() / 1e6
	err := json.Unmarshal([]byte(dataStr), &apiMap)
	if err != nil {
		fmt.Printf("json to map error,err:%s", err)
		return apiMap
	}
	err = json.Unmarshal([]byte(payload), &payloadMap)
	if err != nil {
		fmt.Printf("json to map error,err:%s", err)
		return apiMap
	}
	parameters := ""
	for key, value := range payloadMap {
		if urlEncode {
			encodedStr := url.QueryEscape(fmt.Sprintf("%v", value))
			encodedStr = strings.ReplaceAll(encodedStr, "+", "%20")
			parameters = parameters + key + "=" + encodedStr + "&"
		} else {
			parameters = parameters + key + "=" + fmt.Sprintf("%v", value) + "&"
		}
	}
	parameters += "timestamp=" + fmt.Sprintf("%d", timestemp)
	apiMap["parameters"] = fmt.Sprintf("%v", parameters)
	return apiMap
}
