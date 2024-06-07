package mexcsdk

import (
	"fmt"

	"v1/mexcsdk/client"
	"v1/mexcsdk/model"

	"github.com/bytedance/sonic"
)

type MexcFutureClient struct {
	MexcClient *client.MexcBaseClient
}

func (m *MexcFutureClient) Init(api, sec, url string) *MexcFutureClient {
	fmt.Println(url)
	m.MexcClient = new(client.MexcBaseClient).Init(api, sec, url, "mexc_future")
	return m
}

func (m *MexcFutureClient) GetSymbolDetail(params map[string]string) (model.ApiResponse, error) {
	result, err := m.MexcClient.DoGet("/api/v1/contract/detail", params)
	if err != nil {
		return model.ApiResponse{}, err
	}
	fmt.Println(result)
	balance := model.ApiResponse{}
	sonic.Unmarshal([]byte(result), &balance)
	return balance, nil
}
