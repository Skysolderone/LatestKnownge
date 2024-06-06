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
