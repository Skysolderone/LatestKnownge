package bybit

import (
	"context"
	"net/url"
	"strings"
)

type BalanceService struct {
	Service
	coin        []string
	accountType string
}

func (s *BalanceService) Coin(v ...string) *BalanceService {
	s.coin = v
	return s
}

func (s *BalanceService) AccountType(v string) *BalanceService {
	s.accountType = v
	return s
}

type AccountBalance struct {
	TotalEquity            string        `json:"totalEquity"`
	AccountIMRate          string        `json:"accountIMRate"`
	TotalMarginBalance     string        `json:"totalMarginBalance"`
	TotalInitialMargin     string        `json:"totalInitialMargin"`
	AccountType            string        `json:"accountType"`
	TotalAvailableBalance  string        `json:"totalAvailableBalance"`
	AccountMMRate          string        `json:"accountMMRate"`
	TotalPerpUPL           string        `json:"totalPerpUPL"`
	TotalWalletBalance     string        `json:"totalWalletBalance"`
	AccountLTV             string        `json:"accountLTV"`
	TotalMaintenanceMargin string        `json:"totalMaintenanceMargin"`
	Coin                   []CoinBalance `json:"coin"`
}

type CoinBalance struct {
	AvailableToBorrow   string `json:"availableToBorrow"`
	Bonus               string `json:"bonus"`
	Free                string `json:"free"`
	AccruedInterest     string `json:"accruedInterest"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	TotalOrderIM        string `json:"totalOrderIM"`
	Equity              string `json:"equity"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UsdValue            string `json:"usdValue"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	BorrowAmount        string `json:"borrowAmount"`
	TotalPositionIM     string `json:"totalPositionIM"`
	WalletBalance       string `json:"walletBalance"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Locked              string `json:"locked"`
	MarginCollateral    bool   `json:"marginCollateral"`
	Coin                string `json:"coin"`
}

type BalanceResponse struct {
	List []AccountBalance `json:"list"`
}

func (s *BalanceService) Do(ctx context.Context) ([]AccountBalance, error) {
	params := url.Values{}
	if len(s.coin) > 0 {
		params.Add("coin", strings.Join(s.coin, ","))
	}
	if s.accountType != "" {
		params.Add("accountType", s.accountType)
	}
	s.setQuery(params)

	resp := BalanceResponse{
		List: make([]AccountBalance, 0),
	}
	_, err := s.doAuthRequest(ctx, "GET", "/v5/account/wallet-balance", &resp)
	if err != nil {
		return nil, err
	}
	// if len(resp) == 0 {
	// 	return nil, ErrNilData
	// }
	return resp.List, nil
}

// func (b *BalanceResponse) Asset() []exchange.Asset {
// 	res := make([]exchange.Asset, len(b.Details))
// 	var usdtIndex int
// 	for i, r := range b.Details {
// 		res[i].Symbol = r.Asset
// 		res[i].Free = r.AvailEq
// 		if r.Asset == "USDT" {
// 			usdtIndex = i
// 		} else {
// 			res[i].Balance = decimal.RequireFromString(r.Eq)
// 			res[i].Valaution = decimal.RequireFromString(r.EqUsd)
// 		}
// 	}
// 	res[usdtIndex].Balance = decimal.RequireFromString(b.TotalEq)
// 	res[usdtIndex].Valaution = res[usdtIndex].Balance
// 	return res
// }

type UpgradeUTAService struct {
	Service
}

type UpgradeUTAResponse struct {
	UnifiedUpdateStatus string `json:"unifiedUpdateStatus"`
	UnifiedUpdateMsg    struct {
		Msg []string `json:"msg"`
	} `json:"unifiedUpdateMsg"`
}

func (s *UpgradeUTAService) Do(ctx context.Context) (resp UpgradeUTAResponse, err error) {
	_, err = s.doAuthRequest(ctx, "POST", "/v5/account/upgrade-to-uta", &resp)
	return
}

type ConfigService struct {
	Service
}

type ConfigResponse struct {
	MarginMode  string `json:"marginMode"`
	UpdatedTime string `json:"updatedTime"`
	// 1 普通帳戶
	// 2 已升級到了統一保證金帳戶，僅支持交易期貨和期權
	// 3 已升級到了統一帳戶，支持交易期貨、期權和現貨
	// 4 UTA Pro，統一帳戶的Pro版本
	UnifiedMarginStatus int    `json:"unifiedMarginStatus"`
	DcpStatus           string `json:"dcpStatus"`
	TimeWindow          int    `json:"timeWindow"`
	SmpGroup            int    `json:"smpGroup"`
	IsMasterTrader      bool   `json:"isMasterTrader"`
}

func (s *Service) Do(ctx context.Context) (*ConfigResponse, error) {
	resp := new(ConfigResponse)
	_, err := s.doAuthRequest(ctx, "GET", "/v5/account/info", &resp)
	if err != nil {
		return nil, err
	}
	// 需要升级为统一账户
	// resp.UnifiedMarginStatus == 1
	return resp, nil
}
