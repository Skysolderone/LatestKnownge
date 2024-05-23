package account

import (
	"fmt"

	"bingx/bingx/account/model"
	"bingx/bingx/common"

	"github.com/bytedance/sonic"
)

type AccountSpotClient struct {
	BingXClient *common.BingXClient
}

func (p *AccountSpotClient) Init() *AccountSpotClient {
	p.BingXClient = new(common.BingXClient).Init()
	return p
}

func (m *AccountSpotClient) GetBanlanceDetail(params map[string]string) (model.AccountResponse, error) {
	result, err := m.BingXClient.DoGet("/openApi/spot/v1/account/balance", params)
	if err != nil {
		fmt.Println(err)
		return model.AccountResponse{}, err
	}
	resp := model.AccountResponse{}
	sonic.Unmarshal([]byte(result), &resp)
	return resp, nil
}
