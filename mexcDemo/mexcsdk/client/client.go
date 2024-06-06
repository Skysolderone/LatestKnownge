package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"v1/mexcsdk/config"
	"v1/mexcsdk/utils"

	"github.com/go-resty/resty/v2"
)

type MexcBaseClient struct {
	Apikey string
	Secret string
	URL    string
	Client *resty.Client
	Ex     string
}

func (m *MexcBaseClient) Init(api, sec, url, ex string) *MexcBaseClient {
	m.URL = config.URL
	if url == "" {
		m.URL = config.URL
	}
	client := resty.New()
	m.Client = client
	m.Apikey = api
	m.Secret = sec
	m.Ex = ex
	return m
}

func (m *MexcBaseClient) DoGet(uri string, query map[string]string) (string, error) {
	body := utils.BuildGetParams(query)
	timestemp := time.Now().UnixNano() / 1e6

	if body == "" {
		body += "timestamp=" + fmt.Sprintf("%d", timestemp)
	} else {
		body += "&timestamp=" + fmt.Sprintf("%d", timestemp)
	}

	sign := computeHmac256(body, m.Secret)
	body += fmt.Sprintf("&signature=%s", sign)
	url := m.URL + uri
	m.Client.Debug = true
	resp, err := m.Client.R().SetHeader("x-ex", m.Ex).SetHeader("X-MEXC-APIKEY", m.Apikey).SetQueryString(body).Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(resp.Body()), nil
}

func (m *MexcBaseClient) DoPost(uri string, query string) (string, error) {
	return "", nil
}

func computeHmac256(strMessage string, strSecret string) string {
	key := []byte(strSecret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strMessage))
	return hex.EncodeToString(h.Sum(nil))
}
