package mexcsdk

import "v1/mexcsdk/client"

type MexcFutureClient struct {
	MexcClient *client.MexcBaseClient
}

func (m *MexcFutureClient) Init(api, sec, url string) *MexcFutureClient {
	m.MexcClient = new(client.MexcBaseClient).Init(api, sec, url, "mexc_future")
	return m
}

func (m *MexcFutureClient) Account() {}
