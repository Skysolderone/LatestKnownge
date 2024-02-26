package bybit

import (
	"context"
	"net/url"
	"strconv"
)

type PositionService struct {
	Service
	category   string
	pair       string
	baseCoin   string
	settleCoin string
	limit      int
	cursot     *int
}

func (s *PositionService) Category(c string) *PositionService {
	s.category = c
	return s
}

func (s *PositionService) Pair(p string) *PositionService {
	s.pair = p
	return s
}

func (s *PositionService) BaseCoin(v string) *PositionService {
	s.baseCoin = v
	return s
}

func (s *PositionService) SettleCoin(v string) *PositionService {
	s.settleCoin = v
	return s
}

func (s *PositionService) Limit(v int) *PositionService {
	s.limit = v
	return s
}

func (s *PositionService) Cursot(v int) *PositionService {
	s.cursot = &v
	return s
}

type Position struct {
	Symbol           string `json:"symbol"`
	Leverage         string `json:"leverage"`
	AvgPrice         string `json:"avgPrice"`
	LiqPrice         string `json:"liqPrice"`
	RiskLimitValue   string `json:"riskLimitValue"`
	TakeProfit       string `json:"takeProfit"`
	PositionValue    string `json:"positionValue"`
	TpslMode         string `json:"tpslMode"`
	RiskId           int    `json:"riskId"`
	TrailingStop     string `json:"trailingStop"`
	UnrealisedPnl    string `json:"unrealisedPnl"`
	MarkPrice        string `json:"markPrice"`
	CumRealisedPnl   string `json:"cumRealisedPnl"`
	PositionMM       string `json:"positionMM"`
	CreatedTime      string `json:"createdTime"`
	PositionIdx      int    `json:"positionIdx"`
	PositionIM       string `json:"positionIM"`
	UpdatedTime      string `json:"updatedTime"`
	Side             string `json:"side"`
	BustPrice        string `json:"bustPrice"`
	Size             string `json:"size"`
	PositionStatus   string `json:"positionStatus"`
	StopLoss         string `json:"stopLoss"`
	TradeMode        int    `json:"tradeMode"`
	AdlRankIndicator int    `json:"adlRankIndicator"`
}

type PositionResponse struct {
	NextPageCursor string    `json:"nextPageCursor"`
	Category       string    `json:"category"`
	List           Positions `json:"list"`
}

func (s *PositionService) Do(ctx context.Context) (*PositionResponse, error) {
	params := url.Values{}
	params.Add("category", s.category)
	if s.pair != "" {
		params.Add("symbol", s.pair)
	}
	if s.baseCoin != "" {
		params.Add("baseCoin", s.baseCoin)
	}
	if s.settleCoin != "" {
		params.Add("settleCoin", s.settleCoin)
	}
	if s.limit > 0 {
		params.Add("limit", strconv.Itoa(s.limit))
	}
	if s.cursot != nil {
		params.Add("cursot", strconv.Itoa(*s.cursot))
	}
	s.setQuery(params)

	resp := new(PositionResponse)
	_, err := s.doAuthRequest(ctx, "GET", "/v5/position/list", resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Positions []Position

// type ValueFunc func(key string) decimal.Decimal

// func (p Positions) Position(value ValueFunc) []exchange.Position {
// 	res := make([]exchange.Position, len(p))
// 	for i, pr := range p {
// 		// _, res[i].Pair = ParseInstId(pr.InstId)
// 		res[i].Key = pr.Symbol
// 		res[i].Leverage = pr.Leverage
// 		res[i].PosSide = exchange.PosSide(strings.ToUpper(pr.PosSide))
// 		res[i].Price = pr.AvgPrice
// 		res[i].Quantity = decimal.RequireFromString(pr.Pos).Mul(value(pr.InstId))
// 	}
// 	return res
// }

// func (p PositionResponses) PositionPair(value ValueFunc) []exchange.PositionPair {
// 	pps := make(map[string]*exchange.PositionPair, len(p))
// 	for _, pr := range p {
// 		p, ok := pps[pr.InstId]
// 		if !ok {
// 			p = &exchange.PositionPair{
// 				Leverage: pr.Lever,
// 				Key:      pr.InstId,
// 			}
// 			_, p.Pair = ParseInstId(pr.InstId)
// 			pps[pr.InstId] = p
// 		}
// 		if exchange.PosSide(strings.ToUpper(pr.PosSide)) == exchange.PosSideLong {
// 			p.Long.Price = pr.AvgPx
// 			p.Long.Quantity = decimal.RequireFromString(pr.Pos).Mul(value(pr.InstId))
// 		} else {
// 			p.Short.Price = pr.AvgPx
// 			p.Short.Quantity = decimal.RequireFromString(pr.Pos).Mul(value(pr.InstId))
// 		}
// 	}
// 	res := make([]exchange.PositionPair, 0, len(pps))
// 	for _, pp := range pps {
// 		res = append(res, *pp)
// 	}
// 	return res
// }

type ChangePositionModeService struct {
	Service

	category   string
	pair       string
	settleCoin string
	mode       int
}

func (s *ChangePositionModeService) Category(v string) *ChangePositionModeService {
	s.category = v
	return s
}

func (s *ChangePositionModeService) Pair(v string) *ChangePositionModeService {
	s.pair = v
	return s
}

func (s *ChangePositionModeService) SettleCoin(v string) *ChangePositionModeService {
	s.settleCoin = v
	return s
}

func (s *ChangePositionModeService) Mode(dualSide bool) *ChangePositionModeService {
	if dualSide {
		s.mode = 3
	} else {
		s.mode = 0
	}
	return s
}

type ChangePositionModeResponse struct {
	PosMode string `json:"posMode"`
}

func (s *ChangePositionModeService) Do(ctx context.Context) error {
	params := make(map[string]any, 16)
	params["category"] = s.category
	params["mode"] = s.mode
	if s.pair != "" {
		params["symbol"] = s.pair
	}
	if s.settleCoin != "" {
		params["coin"] = s.settleCoin
	}
	s.setBody(params)
	_, err := s.doAuthRequest(ctx, "POST", "/v5/position/switch-mode", nil)
	if err != nil {
		return err
	}
	return nil
}

// 统一账户不能设置这个
type ChangeMarginModeService struct {
	Service
	category  string
	pair      string
	tradeMode int
	leverage  string
}

func (s *ChangeMarginModeService) Category(v string) *ChangeMarginModeService {
	s.category = v
	return s
}

func (s *ChangeMarginModeService) Pair(v string) *ChangeMarginModeService {
	s.pair = v
	return s
}

func (s *ChangeMarginModeService) TradeMode(cross bool) *ChangeMarginModeService {
	if cross {
		s.tradeMode = 0
	} else {
		s.tradeMode = 1
	}
	return s
}

func (s *ChangeMarginModeService) Leverage(v uint16) *ChangeMarginModeService {
	s.leverage = strconv.Itoa(int(v))
	return s
}

func (s *ChangeMarginModeService) Do(ctx context.Context) error {
	params := make(map[string]any, 16)
	params["category"] = s.category
	params["tradeMode"] = s.tradeMode
	params["symbol"] = s.pair
	params["buyLeverage"] = s.leverage
	params["sellLeverage"] = s.leverage
	s.setBody(params)
	_, err := s.doAuthRequest(ctx, "POST", "/v5/position/switch-isolated", nil)
	if err != nil {
		return err
	}
	return nil
}

type ChangeLeverageService struct {
	Service
	category string
	pair     string
	lever    string // 杠杆倍数
}

func (s *ChangeLeverageService) Category(v string) *ChangeLeverageService {
	s.category = v
	return s
}

func (s *ChangeLeverageService) Pair(p string) *ChangeLeverageService {
	s.pair = p
	return s
}

func (s *ChangeLeverageService) Leverage(v uint16) *ChangeLeverageService {
	s.lever = strconv.Itoa(int(v))
	return s
}

func (s *ChangeLeverageService) Do(ctx context.Context) error {
	params := make(map[string]any, 16)
	params["buyLeverage"] = s.lever
	params["sellLeverage"] = s.lever
	params["category"] = s.category
	if s.pair != "" {
		params["symbol"] = s.pair
	}
	s.setBody(params)

	_, err := s.doAuthRequest(ctx, "POST", "/v5/position/set-leverage", nil)
	if err != nil {
		if ChangedErr(err) {
			return nil
		}
		return err
	}
	return nil
}
