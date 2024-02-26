package bybit

import (
	"context"
	"net/url"
	"strconv"
)

type InfoService struct {
	Service
	category string
	pair     string
	baseCoin string
	status   string
	limit    int64
	cursor   int64
}

func (s *InfoService) Category(c string) *InfoService {
	s.category = c
	return s
}

func (s *InfoService) Pair(v string) *InfoService {
	s.pair = v
	return s
}

func (s *InfoService) BaseCoin(v string) *InfoService {
	s.baseCoin = v
	return s
}

func (s *InfoService) Status(v string) *InfoService {
	s.status = v
	return s
}

func (s *InfoService) Limit(v int64) *InfoService {
	s.limit = v
	return s
}

func (s *InfoService) Cursor(v int64) *InfoService {
	s.cursor = v
	return s
}

type InfoResponse struct {
	Category       string     `json:"category"`
	List           []PairInfo `json:"list"`
	NextPageCursor string     `json:"nextPageCursor"`
}

type PairInfo struct {
	Symbol          string `json:"symbol"`
	ContractType    string `json:"contractType"`
	Status          string `json:"status"`
	BaseCoin        string `json:"baseCoin"`
	QuoteCoin       string `json:"quoteCoin"`
	LaunchTime      int64  `json:"launchTime,string"`
	DeliveryTime    int64  `json:"deliveryTime,string"`
	DeliveryFeeRate string `json:"deliveryFeeRate"`
	PriceScale      string `json:"priceScale"`
	LeverageFilter  struct {
		MinLeverage  string `json:"minLeverage"`
		MaxLeverage  string `json:"maxLeverage"`
		LeverageStep string `json:"leverageStep"`
	} `json:"leverageFilter"`
	PriceFilter struct {
		MinPrice string `json:"minPrice"`
		MaxPrice string `json:"maxPrice"`
		TickSize string `json:"tickSize"`
	} `json:"priceFilter"`
	LotSizeFilter struct {
		MaxOrderQty         string `json:"maxOrderQty"`
		MinOrderQty         string `json:"minOrderQty"`
		BasePrecision       string `json:"basePrecision"`
		QuotePrecision      string `json:"quotePrecision"`
		QtyStep             string `json:"qtyStep"`
		PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
	} `json:"lotSizeFilter"`
	UnifiedMarginTrade bool   `json:"unifiedMarginTrade"`
	FundingInterval    int64  `json:"fundingInterval"`
	SettleCoin         string `json:"settleCoin"`
}

func (s *InfoService) Do(ctx context.Context) (*InfoResponse, error) {
	params := url.Values{}
	params.Add("category", s.category)
	if s.pair != "" {
		params.Add("symbol", s.pair)
	}
	if s.baseCoin != "" {
		params.Add("baseCoin", s.baseCoin)
	}
	if s.status != "" {
		params.Add("status", s.status)
	}
	if s.limit > 0 {
		params.Add("limit", strconv.FormatInt(s.limit, 10))
	}
	if s.cursor > 0 {
		params.Add("cursor", strconv.FormatInt(s.cursor, 10))
	}

	s.setQuery(params)
	resp := new(InfoResponse)
	err := s.doRequest(ctx, "GET", "/v5/market/instruments-info", &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// type InfoResponses []InfoResponse

// func (r InfoResponses) Info() map[string]exchange.PairInfo {
// 	res := make(map[string]exchange.PairInfo, len(r))
// 	for _, v := range r {
// 		if v.State != "live" {
// 			continue
// 		}
// 		c, p := ParseInstId(v.InstId)
// 		var leverage int
// 		var value decimal.Decimal
// 		if c == exchange.Futures {
// 			value = decimal.RequireFromString(v.Value)
// 			leverage, _ = strconv.Atoi(v.Leverage)
// 		}
// 		res[p.String()] = exchange.PairInfo{
// 			Category:    c,
// 			Pair:        p,
// 			Step:        decimal.RequireFromString(v.LotSize),
// 			Value:       value,
// 			MinSize:     decimal.RequireFromString(v.MinSize),
// 			MaxLeverage: uint16(leverage),
// 		}
// 	}
// 	return res
// }
