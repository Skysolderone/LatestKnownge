package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	HOST       = "open-api.bingx.com"
	API_KEY    = "vVQWDndnEPwM5c08ttGx9zy8FOILsmNO2p67hyUIeCoF1iluEzXva5raa5Hc5RknHtDkdoMMqBwVaakhA"
	API_SECRET = "RS48Ul6r8VzEMakjaQCgYN2fOYJTbhk2GIdOxW24kywfSMuXuAXe2G7XncssGeVGsfRwiJWEJTaOqcmglsQ"
)

func main() {
	dataStr := `{
    "uri": "/openApi/swap/v2/trade/order/test",
    "method": "POST",
    "protocol": "https"
}`
	payload := `{
    "symbol": "BTC-USDT",
    "side": "BUY",
    "positionSide": "LONG",
    "type": "MARKET",
    "quantity": 5
}`
	TIMESTAMP := time.Now().UnixNano() / 1e6
	apiMap := getParameters(dataStr, payload, false, TIMESTAMP)
	sign := computeHmac256(fmt.Sprintf("%v", apiMap["parameters"]), API_SECRET)
	fmt.Println("parameters:", fmt.Sprintf("%v", apiMap["parameters"]))
	fmt.Println("sign:", sign)
	parameters := ""
	contains := strings.ContainsAny(fmt.Sprintf("%v", apiMap["parameters"]), "[{")
	if contains {
		apiMap2 := getParameters(dataStr, payload, true, TIMESTAMP)
		parameters = fmt.Sprintf("%v&signature=%s", apiMap2["parameters"], sign)
	} else {
		parameters = fmt.Sprintf("%v&signature=%s", apiMap["parameters"], sign)
	}
	// fmt.Println(parameters)
	url := fmt.Sprintf("%v://%s%v?%s", apiMap["protocol"], HOST, apiMap["uri"], parameters)
	method := fmt.Sprintf("%v", apiMap["method"])
	client := &http.Client{}
	fmt.Println("url:", url)
	fmt.Println("method:", method)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-BX-APIKEY", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))
	return hex.EncodeToString(h.Sum(nil))
}

func getParameters(dataStr string, payload string, urlEncode bool, timestemp int64) map[string]interface{} {
	var apiMap map[string]interface{}
	var payloadMap map[string]interface{}
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
