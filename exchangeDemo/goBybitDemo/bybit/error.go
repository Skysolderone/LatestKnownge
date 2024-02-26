package bybit

import "fmt"

type Error struct {
	Code int    `json:"retCode"`
	Msg  string `json:"retMsg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("<APIError> retCode=%d, retMsg=%s", e.Code, e.Msg)
}

type WsErr struct {
	Success bool   `json:"success"`
	Code    int    `json:"ret_code"`
	Msg     string `json:"ret_msg"`
}

func (e WsErr) Error() string {
	return fmt.Sprintf("<APIError> retCode=%d, retMsg=%s", e.Code, e.Msg)
}

func (e Error) GetErrCode() int {
	return e.Code
}
