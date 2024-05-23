package common

import (
	"fmt"
	"strconv"
	"time"

	"bingx/bingx/config"
	"bingx/bingx/utils"

	resty "github.com/go-resty/resty/v2"
)

type BingXClient struct {
	ApiKey       string
	ApiSecretKey string
	BaseUrl      string
	HttpClient   *resty.Client
}

func (bc *BingXClient) Init() *BingXClient {
	bc.ApiKey = config.API_KEY
	bc.ApiSecretKey = config.API_SECRET
	bc.BaseUrl = config.HOST
	client := resty.New()
	bc.HttpClient = client
	return bc
}

func (bc *BingXClient) DoGet(uri string, params map[string]string) (string, error) {
	body := utils.BuildGetParams(params)

	timestemp := time.Now().UnixNano() / 1e6
	body += "&timestamp=" + fmt.Sprintf("%d", timestemp)

	url := bc.BaseUrl + uri
	resp, err := bc.HttpClient.R().SetHeader("X-BX-APIKEY", bc.ApiKey).SetQueryString(body).Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(resp.Body()), nil
	// return "", nil
}

func (bc *BingXClient) DoPost(uri string, params map[string]string) (string, error) {
	// body := utils.BuildGetParams(params)

	timestemp := time.Now().UnixNano() / 1e6
	// body += "&timestamp=" + fmt.Sprintf("%d", timestemp)

	url := bc.BaseUrl + uri
	params["timestamp"] = strconv.FormatInt(timestemp, 10)
	resp, err := bc.HttpClient.R().SetHeader("X-BX-APIKEY", bc.ApiKey).SetBody(params).Post(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(resp.Body()), nil
	// return "", nil
}
