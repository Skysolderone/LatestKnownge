package model

type AccountResponse struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	DebugMsg string      `json:"debugMsg"`
	Data     AccountResp `json:"data"`
}
type AccountResp struct {
	Balances []BalanceResp `json:"balances"`
}
type BalanceResp struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// future
type FutureAccountResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data BalanceInfo `json:"balance"`
}

type BalanceInfo struct {
	UserID           string `json:"userId"`
	Asset            string `json:"asset"`
	Balance          string `json:"balance"`
	Equity           string `json:"equity"`
	UnrealizedProfit string `json:"unrealizedProfit"`
	RealisedProfit   string `json:"realisedProfit"`
	AvailableMargin  string `json:"availableMargin"`
	UsedMargin       string `json:"usedMargin"`
	FreezedMargin    string `json:"freezedMargin"`
}

type PositionData struct {
	PositionID       string  `json:"positionId"`
	Symbol           string  `json:"symbol"`
	Currency         string  `json:"currency"`
	PositionAmt      string  `json:"positionAmt"`
	AvailableAmt     string  `json:"availableAmt"`
	PositionSide     string  `json:"positionSide"`
	Isolated         bool    `json:"isolated"`
	AvgPrice         string  `json:"avgPrice"`
	InitialMargin    string  `json:"initialMargin"`
	Leverage         int     `json:"leverage"`
	UnrealizedProfit string  `json:"unrealizedProfit"`
	RealisedProfit   string  `json:"realisedProfit"`
	LiquidationPrice float64 `json:"liquidationPrice"`
}

type PositionResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data []PositionData `json:"data"`
}
