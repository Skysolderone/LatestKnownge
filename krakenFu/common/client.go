package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"v1/config"
	"v1/utils"

	resty "github.com/go-resty/resty/v2"
)

type KrakenClient struct {
	ApiKey       string
	ApiSecretKey string
	BaseUrl      string
	HttpClient   *resty.Client
}

func (bc *KrakenClient) Init(api, sec string) *KrakenClient {
	bc.ApiKey = api
	bc.ApiSecretKey = sec
	bc.BaseUrl = config.HOST
	client := resty.New()
	bc.HttpClient = client
	return bc
}

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

func (bc *KrakenClient) DoGet(uri string, params map[string]string) ([]byte, error) {
	body := utils.BuildGetParams(params)

	// if body == "" {
	// 	body += "recvWindow=60000&timestamp=" + fmt.Sprintf("%d", timestemp)
	// } else {
	// 	body += "&recvWindow=60000&timestamp=" + fmt.Sprintf("%d", timestemp)
	// }
	// nonce := strconv.FormatInt(time.Now().UnixMilli(), 10)
	// payload := body + uri
	// fmt.Println(payload)
	url := bc.BaseUrl + "/derivatives" + uri
	bc.HttpClient.Debug = true

	auth, err := Signature(body, "", uri, bc.ApiSecretKey)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(auth)
	resp, err := bc.HttpClient.R().SetHeader("x-ex", "kf").SetHeader("Content-Type", "application/json").
		SetHeader("APIKey", bc.ApiKey).
		SetHeader("Authent", auth).SetQueryString(body).Get(url)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	return resp.Body(), nil
}

func (bc *KrakenClient) DoPost(uri string, params map[string]string) (string, error) {
	body := utils.BuildGetParams(params)

	timestemp := time.Now().UnixNano() / 1e6
	if body == "" {
		body += "timestamp=" + fmt.Sprintf("%d", timestemp)
	} else {
		body += "&timestamp=" + fmt.Sprintf("%d", timestemp)
	}

	fmt.Println(body)

	sign := computeHmac256(body, bc.ApiSecretKey)
	fmt.Println(sign)
	body += fmt.Sprintf("&signature=%s", sign)
	url := bc.BaseUrl + uri + "?" + body
	// url := bc.BaseUrl + uri
	// bc.HttpClient.Debug = true
	// fmt.Println(url)

	// resp, err := bc.HttpClient.R().
	// 	SetHeader("User-Agent", "").
	// 	SetHeader("X-BX-APIKEY", bc.ApiKey).
	// 	SetHeader("x-ex", "bx").
	// 	SetQueryString(body).
	// 	Post(url)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }
	// return string(resp.Body()), nil
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	req.Header.Add("X-BX-APIKEY", bc.ApiKey)
	req.Header.Add("x-ex", "bx")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// fmt.Println(err)
		return "", err
	}
	return string(resp), nil
}

func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))
	return hex.EncodeToString(h.Sum(nil))
}
