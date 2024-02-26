package bybit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AuthOpts struct {
	Key        string
	Secret     string
	Passphrase string
}

type Service struct {
	c         *Client
	apikey    string
	secretkey string

	query url.Values
	body  any
}

func (s *Service) Auth(apikey, secretkey string) {
	s.apikey = apikey
	s.secretkey = secretkey
}

func (s *Service) setQuery(params url.Values) {
	s.query = params
}

func (s *Service) setBody(data any) {
	s.body = data
}

func (s *Service) doAuthRequest(ctx context.Context, method, reqURI string, respBody any) ([]byte, error) {
	if s.apikey == "" {
		return nil, errors.New("missing apikey")
	}
	var (
		paramStr string
		body     []byte
	)
	switch method {
	case "GET":
		query := s.query.Encode()
		if query != "" {
			reqURI += "?" + s.query.Encode()
		}
		paramStr = query
	case "POST":
		// params.Set("tag", "86d4a3bf87bcBCDE")
		bs, err := json.Marshal(s.body)
		if err != nil {
			return nil, err
		}
		body = bs
		paramStr = string(bs)
	}

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10) //iso time style
	recvWindow := "8000"
	sign := SignParam(timestamp, recvWindow, paramStr, s.apikey, s.secretkey)
	// logger.Debugf("[DoAuthRequest] sign base64: %s, timestamp: %s", signStr, timestamp)
	reqURL := s.c.baseURL + reqURI
	rb, err := doRequest(ctx, method, reqURL, body,
		func(r *http.Request) {
			r.Header = s.c.header.Clone()
			r.Header.Set("X-BAPI-API-KEY", s.apikey)
			r.Header.Set("X-BAPI-TIMESTAMP", timestamp)
			r.Header.Set("X-BAPI-SIGN", sign)
			r.Header.Set("X-BAPI-RECV-WINDOW", recvWindow) // 默认值5000
		},
	)
	if err != nil {
		return nil, err
	}
	if s.c.Debug {
		log.Printf("request: %s, response body: %s", reqURL, string(rb))
	}

	var baseResp BaseResp
	err = json.Unmarshal(rb, &baseResp)
	if err != nil {
		return rb, err
	}

	if baseResp.Code != 0 {
		return baseResp.Result, baseResp.Error
	}

	if respBody == nil {
		return baseResp.Result, nil
	}
	return baseResp.Result, json.Unmarshal(baseResp.Result, respBody)
}

type BaseResp struct {
	Error
	Time   int64           `json:"time"`
	Result json.RawMessage `json:"result"`
	Extra  map[string]any  `json:"retExtInfo"`
}

func (s *Service) doRequest(ctx context.Context, method, reqURI string, respBody any) error {
	var reqBody []byte
	switch method {
	case "GET":
		query := s.query.Encode()
		if query != "" {
			reqURI += "?" + s.query.Encode()
		}
	case "POST":
		// params.Set("tag", "86d4a3bf87bcBCDE")
		bs, err := json.Marshal(s.body)
		if err != nil {
			return err
		}
		reqBody = bs
	}

	reqURL := s.c.baseURL + reqURI
	// fmt.Println("request url", reqURL)
	rb, err := doRequest(ctx, method, reqURL, reqBody, func(r *http.Request) {
		r.Header = s.c.header.Clone()
	})
	if err != nil {
		return err
	}
	if s.c.Debug {
		log.Printf("request: %s, response body: %s", reqURL, string(rb))
	}

	var baseResp BaseResp
	err = json.Unmarshal(rb, &baseResp)
	if err != nil {
		return err
	}

	if baseResp.Code != 0 {
		return errors.New(baseResp.Msg)
	}

	return json.Unmarshal(baseResp.Result, respBody)
}

type RequestOption func(r *http.Request)

func SetHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Add(key, value)
	}
}

func doRequest(ctx context.Context, method, reqUrl string, body []byte, opts ...RequestOption) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, reqUrl, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	for _, opt := range opts {
		opt(req)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// bs, _ := httputil.DumpResponse(resp, true)
	// fmt.Println(string(bs))
	// if resp.StatusCode != 200 {
	// 	return nil, errors.New(resp.Status)
	// }
	return io.ReadAll(resp.Body)
}
