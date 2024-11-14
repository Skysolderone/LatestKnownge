package model

type TradeResponse struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	DebugMsg string    `json:"debugMsg"`
	Data     TradeResp `json:"data"`
}

type TradeResp struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}
type OrderResponse struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	DebugMsg string    `json:"debugMsg"`
	Data     OrderInfo `json:"data"`
}
type OrderInfo struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	Price               string `json:"price"`
	ClientOrderID       string `json:"clientOrderID"`
	StopPrice           string `json:"stopPrice"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	Time                int64  `json:"time"`
	UpdateTime          int64  `json:"updateTime"`
	OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
	Fee                 string `json:"fee"`
	FeeAsset            string `json:"feeAsset"`
}

// future
type FutureTradeResp struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data FuturePlaceOrder `json:"data"`
}
type FuturePlaceOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Type          string `json:"type"`
	ClientOrderID string `json:"clientOrderID"`
	WorkingType   string `json:"workingType"`
}

type FutureOrderResponse struct {
	Code int `json:"code"`
	Data struct {
		Order struct {
			AdvanceAttr   int    `json:"advanceAttr"`
			ReduceOnly    string `json:"reduceOnly"`
			AvgPrice      string `json:"avgPrice"`
			ClientOrderID string `json:"clientOrderId"`
			Commission    string `json:"commission"`
			CumQuote      string `json:"cumQuote"`
			ExecutedQty   string `json:"executedQty"`
			Leverage      string `json:"leverage"`
			OrderID       int64  `json:"orderId"`
			OrderType     string `json:"orderType"`
			OrigQty       string `json:"origQty"`
			PositionID    int    `json:"positionID"`
			PositionSide  string `json:"positionSide"`
			Price         string `json:"price"`
			Profit        string `json:"profit"`
			Side          string `json:"side"`
			Status        string `json:"status"`

			StopGuaranteed bool `json:"stopGuaranteed"`
			StopLoss       struct {
				Price       int    `json:"price"`
				Quantity    int    `json:"quantity"`
				StopPrice   int    `json:"stopPrice"`
				Type        string `json:"type"`
				WorkingType string `json:"workingType"`
			} `json:"stopLoss"`
			StopLossEntrustPrice int    `json:"stopLossEntrustPrice"`
			StopPrice            string `json:"stopPrice"`
			Symbol               string `json:"symbol"`
			TakeProfit           struct {
				Price       int    `json:"price"`
				Quantity    int    `json:"quantity"`
				StopPrice   int    `json:"stopPrice"`
				Type        string `json:"type"`
				WorkingType string `json:"workingType"`
			} `json:"takeProfit"`
			TakeProfitEntrustPrice int64  `json:"takeProfitEntrustPrice"`
			Time                   int64  `json:"time"`
			TriggerOrderID         int64  `json:"triggerOrderId"`
			Type                   string `json:"type"`
			UpdateTime             int64  `json:"updateTime"`
			WorkingType            string `json:"workingType"`
		} `json:"order"`
	} `json:"data"`
	Msg string `json:"msg"`
}

// 持仓模式
type Data struct {
	DualSidePosition string `json:"dualSidePosition"`
}

type TradePositionResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

// leverage
type LeverageResp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data LeverageData `json:"data"`
}
type LeverageData struct {
	Leverage int    `json:"leverage"`
	Symbol   string `json:"symbol"`
}

// marginType
type MarginResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
