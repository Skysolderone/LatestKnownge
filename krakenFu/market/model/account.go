package model

import "time"

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
	Result     string    `json:"result"`
	ServerTime time.Time `json:"serverTime"`
	Accounts   struct {
		Cash struct {
			Type     string `json:"type"`
			Balances struct {
				Xbt float64 `json:"xbt"`
				Xrp float64 `json:"xrp"`
			} `json:"balances"`
		} `json:"cash"`
		FiXbtusd struct {
			Type     string `json:"type"`
			Currency string `json:"currency"`
			Balances struct {
				FIXBTUSD171215 int     `json:"FI_XBTUSD_171215"`
				FIXBTUSD180615 int     `json:"FI_XBTUSD_180615"`
				Xbt            float64 `json:"xbt"`
				Xrp            int     `json:"xrp"`
			} `json:"balances"`
			Auxiliary struct {
				Af  float64 `json:"af"`
				Pnl float64 `json:"pnl"`
				Pv  float64 `json:"pv"`
			} `json:"auxiliary"`
			MarginRequirements struct {
				Im float64 `json:"im"`
				Mm float64 `json:"mm"`
				Lt float64 `json:"lt"`
				Tt float64 `json:"tt"`
			} `json:"marginRequirements"`
			TriggerEstimates struct {
				Im int `json:"im"`
				Mm int `json:"mm"`
				Lt int `json:"lt"`
				Tt int `json:"tt"`
			} `json:"triggerEstimates"`
		} `json:"fi_xbtusd"`
		Flex struct {
			Type       string `json:"type"`
			Currencies struct {
				XBT struct {
					Quantity   float64 `json:"quantity"`
					Value      float64 `json:"value"`
					Collateral float64 `json:"collateral"`
					Available  float64 `json:"available"`
				} `json:"XBT"`
				USD struct {
					Quantity   int `json:"quantity"`
					Value      int `json:"value"`
					Collateral int `json:"collateral"`
					Available  int `json:"available"`
				} `json:"USD"`
				EUR struct {
					Quantity   float64 `json:"quantity"`
					Value      float64 `json:"value"`
					Collateral float64 `json:"collateral"`
					Available  float64 `json:"available"`
				} `json:"EUR"`
			} `json:"currencies"`
			BalanceValue            float64 `json:"balanceValue"`
			PortfolioValue          float64 `json:"portfolioValue"`
			CollateralValue         float64 `json:"collateralValue"`
			InitialMargin           int     `json:"initialMargin"`
			InitialMarginWithOrders int     `json:"initialMarginWithOrders"`
			MaintenanceMargin       int     `json:"maintenanceMargin"`
			Pnl                     int     `json:"pnl"`
			UnrealizedFunding       int     `json:"unrealizedFunding"`
			TotalUnrealized         int     `json:"totalUnrealized"`
			TotalUnrealizedAsMargin int     `json:"totalUnrealizedAsMargin"`
			MarginEquity            float64 `json:"marginEquity"`
			AvailableMargin         float64 `json:"availableMargin"`
		} `json:"flex"`
	} `json:"accounts"`
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
	UpdateTime       int64   `json:"updateTime"`
	MarkPrice        string  `json:"markPrice"`
}

type PositionResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data []PositionData `json:"data"`
}
