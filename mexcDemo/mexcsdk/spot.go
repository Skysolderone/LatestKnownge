package mexcsdk

import (
	"fmt"

	"v1/mexcsdk/client"
	"v1/mexcsdk/model"

	"github.com/bytedance/sonic"
)

type MexcSpotClient struct {
	MexcClient *client.MexcBaseClient
}

func (m *MexcSpotClient) Init(api, sec, url string) *MexcSpotClient {
	m.MexcClient = new(client.MexcBaseClient).Init(api, sec, url, "mexc_spot")
	return m
}

func (m *MexcSpotClient) Account(params map[string]string) (model.AccountInfo, error) {
	result, err := m.MexcClient.DoGet("/api/v3/account", params)
	if err != nil {
		return model.AccountInfo{}, err
	}
	balance := model.AccountInfo{}
	sonic.Unmarshal([]byte(result), &balance)
	return balance, nil
}

func (m *MexcSpotClient) GetOrder(params map[string]string) (model.AccountInfo, error) {
	result, err := m.MexcClient.DoGet("/api/v3/order", params)
	if err != nil {
		return model.AccountInfo{}, err
	}
	fmt.Println(result)
	balance := model.AccountInfo{}
	sonic.Unmarshal([]byte(result), &balance)
	return balance, nil
}
