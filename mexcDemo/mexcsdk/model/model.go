package model

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountInfo struct {
	CanTrade    bool      `json:"canTrade"`
	CanWithdraw bool      `json:"canWithdraw"`
	CanDeposit  bool      `json:"canDeposit"`
	UpdateTime  *int64    `json:"updateTime"` // 使用指针类型来处理null值
	AccountType string    `json:"accountType"`
	Balances    []Balance `json:"balances"`
	Permissions []string  `json:"permissions"`
}

type Symbol struct {
	Symbol                     string        `json:"symbol"`
	Status                     string        `json:"status"`
	BaseAsset                  string        `json:"baseAsset"`
	BaseAssetPrecision         int           `json:"baseAssetPrecision"`
	QuoteAsset                 string        `json:"quoteAsset"`
	QuotePrecision             int           `json:"quotePrecision"`
	QuoteAssetPrecision        int           `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int           `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int           `json:"quoteCommissionPrecision"`
	OrderTypes                 []string      `json:"orderTypes"`
	QuoteOrderQtyMarketAllowed bool          `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool          `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool          `json:"isMarginTradingAllowed"`
	QuoteAmountPrecision       string        `json:"quoteAmountPrecision"`
	BaseSizePrecision          string        `json:"baseSizePrecision"`
	Permissions                []string      `json:"permissions"`
	Filters                    []interface{} `json:"filters"`
	MaxQuoteAmount             string        `json:"maxQuoteAmount"`
	MakerCommission            string        `json:"makerCommission"`
	TakerCommission            string        `json:"takerCommission"`
}

type ExchangeInfo struct {
	Timezone   string   `json:"timezone"`
	ServerTime int64    `json:"serverTime"`
	Symbols    []Symbol `json:"symbols"`
}


type SymbolData struct {
	Symbol                    string   `json:"symbol"`
	DisplayName               string   `json:"displayName"`
	DisplayNameEn             string   `json:"displayNameEn"`
	PositionOpenType          int      `json:"positionOpenType"`
	BaseCoin                  string   `json:"baseCoin"`
	QuoteCoin                 string   `json:"quoteCoin"`
	SettleCoin                string   `json:"settleCoin"`
	ContractSize              float64  `json:"contractSize"`
	MinLeverage               int      `json:"minLeverage"`
	MaxLeverage               int      `json:"maxLeverage"`
	PriceScale                int      `json:"priceScale"`
	VolScale                  int      `json:"volScale"`
	AmountScale               int      `json:"amountScale"`
	PriceUnit                 float64  `json:"priceUnit"`
	VolUnit                   int      `json:"volUnit"`
	MinVol                    int      `json:"minVol"`
	MaxVol                    int      `json:"maxVol"`
	BidLimitPriceRate         float64  `json:"bidLimitPriceRate"`
	AskLimitPriceRate         float64  `json:"askLimitPriceRate"`
	TakerFeeRate              float64  `json:"takerFeeRate"`
	MakerFeeRate              float64  `json:"makerFeeRate"`
	MaintenanceMarginRate     float64  `json:"maintenanceMarginRate"`
	InitialMarginRate         float64  `json:"initialMarginRate"`
	RiskBaseVol               int      `json:"riskBaseVol"`
	RiskIncrVol               int      `json:"riskIncrVol"`
	RiskIncrMmr               float64  `json:"riskIncrMmr"`
	RiskIncrImr               float64  `json:"riskIncrImr"`
	RiskLevelLimit            int      `json:"riskLevelLimit"`
	PriceCoefficientVariation float64  `json:"priceCoefficientVariation"`
	IndexOrigin               []string `json:"indexOrigin"`
	State                     int      `json:"state"`
	IsNew                     bool     `json:"isNew"`
	IsHot                     bool     `json:"isHot"`
	IsHidden                  bool     `json:"isHidden"`
	ConceptPlate              []string `json:"conceptPlate"`
	RiskLimitType             string   `json:"riskLimitType"`
	MaxNumOrders              []int    `json:"maxNumOrders"`
	MarketOrderMaxLevel       int      `json:"marketOrderMaxLevel"`
	MarketOrderPriceLimitRate1 float64 `json:"marketOrderPriceLimitRate1"`
	MarketOrderPriceLimitRate2 float64 `json:"marketOrderPriceLimitRate2"`
	TriggerProtect            float64  `json:"triggerProtect"`
	Appraisal                 int      `json:"appraisal"`
	ShowAppraisalCountdown    int      `json:"showAppraisalCountdown"`
	AutomaticDelivery         int      `json:"automaticDelivery"`
	ApiAllowed                bool     `json:"apiAllowed"`
}

type ApiResponse struct {
	Success bool         `json:"success"`
	Code    int          `json:"code"`
	Data    []SymbolData `json:"data"`
}