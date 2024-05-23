package model

type SymbolResponse struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	DebugMsg string     `json:"debugMsg"`
	Data     SymbolResp `json:"data"`
}
type SymbolResp struct {
	Symbols []SymbolInfo `json:"symbols"`
}

type SymbolInfo struct {
	Symbol       string  `json:"symbol"`
	MinQty       float64 `json:"minQty"`
	MaxQty       float64 `json:"maxQty"`
	MinNotional  float64 `json:"minNotional"`
	MaxNotional  float64 `json:"maxNotional"`
	Status       int     `json:"status"`
	TickSize     float64 `json:"tickSize"`
	StepSize     float64 `json:"stepSize"`
	ApiStateSell bool    `json:"apiStateSell"`
	ApiStateBuy  bool    `json:"apiStateBuy"`
	TimeOnline   int     `json:"timeOnline"`
}
type FutreSymbolResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data []ContractInfo `json:"data"`
}

type ContractInfo struct {
	ContractID        string  `json:"contractId"`
	Symbol            string  `json:"symbol"`
	Size              string  `json:"size"`
	QuantityPrecision int     `json:"quantityPrecision"`
	PricePrecision    int     `json:"pricePrecision"`
	FeeRate           float64 `json:"feeRate"`
	MakerFeeRate      float64 `json:"makerFeeRate"`
	TakerFeeRate      float64 `json:"takerFeeRate"`
	TradeMinLimit     float64 `json:"tradeMinLimit"`
	TradeMinQuantity  float64 `json:"tradeMinQuantity"`
	TradeMinUSDT      float64 `json:"tradeMinUSDT"`
	MaxLongLeverage   int     `json:"maxLongLeverage"`
	MaxShortLeverage  int     `json:"maxShortLeverage"`
	Currency          string  `json:"currency"`
	Asset             string  `json:"asset"`
	Status            int     `json:"status"`
	ApiStateOpen      string  `json:"apiStateOpen"`
	ApiStateClose     string  `json:"apiStateClose"`
	EnsureTrigger     bool    `json:"ensureTrigger"`
	TriggerFeeRate    string  `json:"triggerFeeRate"`
}
