package bybit

import (
	"context"
	"time"
)

type ApiService struct {
	Service
}

type ApiInfo struct {
	ID          string `json:"id"`
	Note        string `json:"note"`
	Apikey      string `json:"apiKey"`
	ReadOnly    int    `json:"readOnly"`
	Secret      string `json:"secret"`
	Permissions struct {
		ContractTrade []string `json:"ContractTrade"`
		Spot          []string `json:"Spot"`
		Wallet        []string `json:"Wallet"`
		Options       []string `json:"Options"`
		Derivatives   []string `json:"Derivatives"`
		CopyTrading   []string `json:"CopyTrading"`
		BlockTrade    []string `json:"BlockTrade"`
		Exchange      []string `json:"Exchange"`
		Nft           []string `json:"NFT"`
		Affiliate     []string `json:"Affiliate"`
	} `json:"permissions"`
	Ips           []string  `json:"ips"`
	Type          int       `json:"type"`
	DeadlineDay   int       `json:"deadlineDay"`
	ExpiredAt     time.Time `json:"expiredAt"`
	CreatedAt     time.Time `json:"createdAt"`
	Unified       int       `json:"unified"`
	Uta           int       `json:"uta"`
	UserId        int       `json:"userID"`
	InviterId     int       `json:"inviterID"`
	VipLevel      string    `json:"vipLevel"`
	MktMakerLevel string    `json:"mktMakerLevel"`
	AffiliateId   int       `json:"affiliateID"`
	RsaPublickey  string    `json:"rsaPublicKey"`
	IsMaster      bool      `json:"isMaster"`
	ParentUid     string    `json:"parentUid"`
	KycLevel      string    `json:"kycLevel"`
	KycRegion     string    `json:"kycRegion"`
}

func (s *ApiService) Do(ctx context.Context) (*ApiInfo, error) {
	resp := new(ApiInfo)
	_, err := s.doAuthRequest(ctx, "GET", "/v5/user/query-api", &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
