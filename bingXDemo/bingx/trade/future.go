package trade

import (
	"fmt"

	"bingx/bingx/common"
	"bingx/bingx/trade/model"

	"github.com/bytedance/sonic"
)

type TradeFutureClient struct {
	BingXClient *common.BingXClient
}

func (p *TradeFutureClient) Init() *TradeFutureClient {
	p.BingXClient = new(common.BingXClient).Init()
	return p
}

func (m *TradeFutureClient) PlaceOrder(params map[string]string) (model.FutureTradeResp, error) {
	// /openApi/swap/v2/trade/order/test  测试下单接口
	result, err := m.BingXClient.DoPost("/openApi/swap/v2/trade/order", params)
	if err != nil {
		fmt.Println(err)
		return model.FutureTradeResp{}, err
	}
	resp := model.FutureTradeResp{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

func (m *TradeFutureClient) GetOrder(params map[string]string) (model.FutureOrderResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/swap/v2/trade/order", params)
	if err != nil {
		fmt.Println(err)
		return model.FutureOrderResponse{}, err
	}
	resp := model.FutureOrderResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 获取持仓模式
func (m *TradeFutureClient) GetPositionMode(params map[string]string) (model.PositionResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/swap/v1/positionSide/dual", params)
	if err != nil {
		fmt.Println(err)
		return model.PositionResponse{}, err
	}
	resp := model.PositionResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 设置持仓模式
func (m *TradeFutureClient) SetPositionMode(params map[string]string) (model.PositionResponse, error) {
	result, err := m.BingXClient.DoPost("/openApi/swap/v1/positionSide/dual", params)
	if err != nil {
		fmt.Println(err)
		return model.PositionResponse{}, err
	}
	resp := model.PositionResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 设置杠杆
func (m *TradeFutureClient) SetLeverage(params map[string]string) (model.LeverageResp, error) {
	result, err := m.BingXClient.DoPost("/openApi/swap/v2/trade/leverage", params)
	if err != nil {
		fmt.Println(err)
		return model.LeverageResp{}, err
	}
	resp := model.LeverageResp{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 更改保证金模式
func (m *TradeFutureClient) SetMarginType(params map[string]string) (model.MarginResp, error) {
	result, err := m.BingXClient.DoPost("/openApi/swap/v2/trade/marginType", params)
	if err != nil {
		fmt.Println(err)
		return model.MarginResp{}, err
	}
	resp := model.MarginResp{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}
