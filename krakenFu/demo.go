package main

import (
	"fmt"

	"v1/market"
)

// import "v1/market"

func main() {
	client := new(market.MarketFutureClient).Init(Api, Sec)
	result, err := client.GetBanlanceDetail(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", result)
}

// import (
// 	"crypto/hmac"
// 	"crypto/sha256"
// 	"crypto/sha512"
// 	"encoding/base64"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	urls := "https://futures.kraken.com/derivatives/api/v3/accounts"
// 	method := "GET"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, urls, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// concatenatedString := postData + nonce + endpointPath

// 	// // Step 2: Hash the result of step 1 with the SHA-256 algorithm
// 	// sha256Hash := sha256.Sum256([]byte(concatenatedString))

// 	// // Step 3: Base64-decode your api_secret
// 	// decodedSecret, err := base64.StdEncoding.DecodeString(apiSecret)
// 	// if err != nil {
// 	// 	fmt.Println("Error decoding API secret:", err)
// 	// 	return
// 	// }

// 	// // Step 4: Hash the result of step 2 with the HMAC-SHA-512 algorithm using the decoded secret
// 	// hmacHash := hmac.New(sha512.New, decodedSecret)
// 	// hmacHash.Write(sha256Hash[:])
// 	// hmacResult := hmacHash.Sum(nil)

// 	// // Step 5: Base64-encode the result of step 4
// 	// authent := base64.StdEncoding.EncodeToString(hmacResult)

// 	nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
// 	fmt.Println(nonce)
// 	params := url.Values{}
// 	// params.Add("data", "string")
// 	// params.Add("sx", "eea")
// 	postdata := params.Encode()
// 	fmt.Println(postdata)
// 	payload := postdata + nonce + "/api/v3/accounts"

// 	sha := sha256.New()
// 	sha.Write([]byte(payload))
// 	shasum := sha.Sum(nil)
// 	s, err := base64.URLEncoding.DecodeString(sec)
// 	if err != nil {
// 		return
// 	}
// 	mac := hmac.New(sha512.New, s)
// 	mac.Write(shasum[:])
// 	macsum := mac.Sum(nil)
// 	auth := base64.StdEncoding.EncodeToString(macsum)
// 	// auth, err := SignParam("/api/v3/accounts", sec, nonce)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	req.Header.Add("Accept", "application/json")
// 	req.Header.Add("APIKey", api)
// 	fmt.Println(auth)
// 	req.Header.Add("Authent", auth)

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("%#v", res)
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(body))
// }
