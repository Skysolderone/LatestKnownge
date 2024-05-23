package market

import (
	"fmt"

	"bingx/bingx/common"
	"bingx/bingx/market/model"

	"github.com/bytedance/sonic"
)

type MarketFutureClient struct {
	BingXClient *common.BingXClient
}

func (p *MarketFutureClient) Init() *MarketFutureClient {
	p.BingXClient = new(common.BingXClient).Init()
	return p
}

func (m *MarketFutureClient) GetSymbolDetail(params map[string]string) (model.FutreSymbolResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/swap/v2/quote/contracts", params)
	if err != nil {
		fmt.Println(err)
		return model.FutreSymbolResponse{}, err
	}
	resp := model.FutreSymbolResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}
