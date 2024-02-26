package bybit

import (
	"context"
	"net/url"
	"strconv"
)

const MaxKlineLimit = 300

type KlineService struct {
	Service

	category string
	pair     string
	interval int64
	end      string
	start    string
	limit    string
}

func (s *KlineService) Category(c string) *KlineService {
	s.category = c
	return s
}

func (s *KlineService) Pair(p string) *KlineService {
	s.pair = p
	return s
}

func (s *KlineService) Start(v int64) *KlineService {
	s.start = strconv.FormatInt(v*1000, 10)
	return s
}

func (s *KlineService) End(v int64) *KlineService {
	s.end = strconv.FormatInt(v*1000, 10)
	return s
}

func (s *KlineService) Limit(v int64) *KlineService {
	s.limit = strconv.FormatInt(v, 10)
	return s
}

func (s *KlineService) Interval(v int64) *KlineService {
	s.interval = v
	return s
}

type KlineResponse struct {
	Symbol   string     `json:"symbol"`
	Category string     `json:"category"`
	List     [][]string `json:"list"`
}

func (s *KlineService) Do(ctx context.Context) ([][]string, error) {
	params := url.Values{}
	if s.category == "" {
		s.category = Spot
	}
	params.Set("category", s.category)
	params.Set("symbol", s.pair)
	if s.start != "" {
		params.Set("start", s.start)
	}
	if s.end != "" {
		params.Set("end", s.end)
	}
	params.Set("interval", strconv.FormatInt(s.interval/60, 10))
	if s.limit != "" {
		params.Set("limit", s.limit)
	}
	s.setQuery(params)

	resp := new(KlineResponse)
	err := s.doRequest(ctx, "GET", "/v5/market/kline", resp)
	if err != nil {
		return nil, err
	}

	res := make([][]string, len(resp.List))
	for _, v := range resp.List {
		ot, _ := strconv.ParseInt(v[0], 10, 64)
		k := []string{
			strconv.FormatInt(ot/1000-s.interval+1, 10),
			v[1],
			v[2],
			v[3],
			v[4],
			v[5],
			v[6],
		}
		res = append(res, k)
	}
	return res, nil
}
