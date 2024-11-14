package market

import (
	"fmt"

	"v1/common"
	"v1/market/model"

	"github.com/bytedance/sonic"
)

type MarketFutureClient struct {
	KrakenClient *common.KrakenClient
}

func (p *MarketFutureClient) Init(api, sec string) *MarketFutureClient {
	p.KrakenClient = new(common.KrakenClient).Init(api, sec)
	return p
}

func (m *MarketFutureClient) GetSymbolDetail(params map[string]string) (model.FutreSymbolResponse, error) {
	result, err := m.KrakenClient.DoGet("/openApi/swap/v2/quote/contracts", params)
	if err != nil {
		fmt.Println(err)
		return model.FutreSymbolResponse{}, err
	}
	resp := model.FutreSymbolResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (m *MarketFutureClient) PlaceOrder(params map[string]string) (model.FutureTradeResp, error) {
	// /openApi/swap/v2/trade/order/test  测试下单接口
	result, err := m.KrakenClient.DoPost("/openApi/swap/v2/trade/order", params)
	if err != nil {
		fmt.Println(err)
		return model.FutureTradeResp{}, err
	}
	resp := model.FutureTradeResp{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (m *MarketFutureClient) GetOrder(params map[string]string) (model.FutureOrderResponse, error) {
	result, err := m.KrakenClient.DoGet("/openApi/swap/v2/trade/order", params)
	if err != nil {
		fmt.Println(err)
		return model.FutureOrderResponse{}, err
	}
	resp := model.FutureOrderResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 获取持仓模式
func (m *MarketFutureClient) GetPositionMode(params map[string]string) (model.TradePositionResponse, error) {
	result, err := m.KrakenClient.DoGet("/openApi/swap/v1/positionSide/dual", params)
	if err != nil {
		fmt.Println(err)
		return model.TradePositionResponse{}, err
	}
	resp := model.TradePositionResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 设置持仓模式
func (m *MarketFutureClient) SetPositionMode(params map[string]string) (model.TradePositionResponse, error) {
	result, err := m.KrakenClient.DoPost("/openApi/swap/v1/positionSide/dual", params)
	if err != nil {
		fmt.Println(err)
		return model.TradePositionResponse{}, err
	}
	resp := model.TradePositionResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 设置杠杆
func (m *MarketFutureClient) SetLeverage(params map[string]string) (model.LeverageResp, error) {
	result, err := m.KrakenClient.DoPost("/openApi/swap/v2/trade/leverage", params)
	if err != nil {
		fmt.Println(err)
		return model.LeverageResp{}, err
	}
	resp := model.LeverageResp{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 更改保证金模式
func (m *MarketFutureClient) SetMarginType(params map[string]string) (model.MarginResp, error) {
	result, err := m.KrakenClient.DoPost("/openApi/swap/v2/trade/marginType", params)
	if err != nil {
		fmt.Println(err)
		return model.MarginResp{}, err
	}
	resp := model.MarginResp{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (m *MarketFutureClient) GetBanlanceDetail(params map[string]string) (model.FutureAccountResponse, error) {
	result, err := m.KrakenClient.DoGet("/api/v3/accounts", params)
	fmt.Printf("%#v", result)
	if err != nil {
		fmt.Println(err)
		return model.FutureAccountResponse{}, err
	}
	resp := model.FutureAccountResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 获取所有持仓
func (m *MarketFutureClient) GetUserPositions(params map[string]string) (model.PositionResponse, error) {
	result, err := m.KrakenClient.DoGet("/openApi/swap/v2/user/positions", params)
	if err != nil {
		fmt.Println(err)
		return model.PositionResponse{}, err
	}
	resp := model.PositionResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}
