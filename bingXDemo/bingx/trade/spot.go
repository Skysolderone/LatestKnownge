package trade

import (
	"fmt"

	"bingx/bingx/common"
	"bingx/bingx/trade/model"

	"github.com/bytedance/sonic"
)

type TradeSpotClient struct {
	BingXClient *common.BingXClient
}

func (p *TradeSpotClient) Init() *TradeSpotClient {
	p.BingXClient = new(common.BingXClient).Init()
	return p
}

func (m *TradeSpotClient) PlaceOrder(params map[string]string) (model.TradeResponse, error) {
	result, err := m.BingXClient.DoPost("/openApi/spot/v1/trade/order", params)
	if err != nil {
		fmt.Println(err)
		return model.TradeResponse{}, err
	}
	resp := model.TradeResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (m *TradeSpotClient) GetOrder(params map[string]string) (model.OrderResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/spot/v1/trade/query", params)
	if err != nil {
		fmt.Println(err)
		return model.OrderResponse{}, err
	}
	resp := model.OrderResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}
