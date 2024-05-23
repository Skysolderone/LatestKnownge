package account

import (
	"fmt"

	"bingx/bingx/account/model"
	"bingx/bingx/common"

	"github.com/bytedance/sonic"
)

type AccountFutureClient struct {
	BingXClient *common.BingXClient
}

func (p *AccountFutureClient) Init() *AccountFutureClient {
	p.BingXClient = new(common.BingXClient).Init()
	return p
}

func (m *AccountFutureClient) GetBanlanceDetail(params map[string]string) (model.FutureAccountResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/swap/v2/user/balance", params)
	if err != nil {
		fmt.Println(err)
		return model.FutureAccountResponse{}, err
	}
	resp := model.FutureAccountResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}

// 获取所有持仓
func (m *AccountFutureClient) GetUserPositions(params map[string]string) (model.PositionResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/swap/v2/user/positions", params)
	if err != nil {
		fmt.Println(err)
		return model.PositionResponse{}, err
	}
	resp := model.PositionResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}
