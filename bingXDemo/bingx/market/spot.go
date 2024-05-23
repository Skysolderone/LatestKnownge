package market

import (
	"fmt"

	"bingx/bingx/common"
	"bingx/bingx/market/model"

	"github.com/bytedance/sonic"
)

type MarketSpotClient struct {
	BingXClient *common.BingXClient
}

func (p *MarketSpotClient) Init() *MarketSpotClient {
	p.BingXClient = new(common.BingXClient).Init()
	return p
}

func (m *MarketSpotClient) GetSymbolDetail(params map[string]string) (model.SymbolResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/spot/v1/common/symbols", params)
	if err != nil {
		fmt.Println(err)
		return model.SymbolResponse{}, err
	}
	resp := model.SymbolResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

