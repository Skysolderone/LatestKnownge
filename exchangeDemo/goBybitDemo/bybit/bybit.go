package bybit

import (
	"errors"
)

// https
const (
	baseURL        = "https://api.bybit.com"
	optionsBaseURL = "https://api.bytick.com"
)

// Websocket
const (
	baseWsURL = "wss://stream.bybit.com"
	// priWsURL  = "wss://stream.bybit.com/v5/private"

	// spotWsURL    = "wss://stream.bybit.com/v5/public/spot"
	// futuresWsURL = "wss://stream.bybit.com/v5/public/linear"

	priWsPath     = "/v5/private"
	spotWsPath    = "/v5/public/spot"
	futuresWsParh = "/v5/public/linear"
)

const (
	Spot    = "spot"
	Futures = "linear"
)

func ChangedErr(err error) bool {
	var apiErr Error
	if errors.As(err, &apiErr) {
		return apiErr.Code == 110043
	}
	return false
}

func ApiInvalid(err error) bool {
	var apiErr Error
	if errors.As(err, &apiErr) {
		switch apiErr.Code {
		case 33004, 10003, 10005:
			return true
		}
	}
	return false
}
