package bybit

import (
	"context"
	"net/url"
	"strconv"
	"time"
	"trading/common"
)

type OrderService struct {
	Service

	category  string
	pair      string
	side      string
	orderType string
	size      string // 委托数量
	price     string // 委托价格

	positionIdx *int // 0: 單向持倉	1: 買側雙向持倉 2: 賣側雙向持倉

	orderLinkId string
	timeInForce string
}

func (s *OrderService) Pair(c, p string) *OrderService {
	s.category = c
	s.pair = p
	return s
}

// cross 逐仓
// isolated 全仓
// cash 非保证金
// func (s *OrderService) Mode(v string) *OrderService {
// 	s.tdMode = v
// 	return s
// }

func (s *OrderService) Side(v string) *OrderService {
	s.side = v
	return s
}

// limit market
func (s *OrderService) Type(v string) *OrderService {
	s.orderType = v
	return s
}

// 若category=spot，且是Market Buy单，则qty表示为报价币种金额
func (s *OrderService) Size(v string) *OrderService {
	s.size = v
	return s
}

func (s *OrderService) Price(v string) *OrderService {
	s.price = v
	return s
}

func (s *OrderService) PosSide(v string) *OrderService {
	var idx int
	switch v {
	case "Short":
		idx = 2
	case "Long":
		idx = 1
	}
	s.positionIdx = &idx
	return s
}

func (s *OrderService) OrderID(v string) *OrderService {
	s.orderLinkId = v
	return s
}

func (s *OrderService) TimeInForce(v string) *OrderService {
	s.timeInForce = v
	return s
}

type CreateOrderResponse struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

// todo:
func (s *OrderService) Do(ctx context.Context) (*CreateOrderResponse, error) {
	params := make(map[string]any, 16)
	params["symbol"] = s.pair
	params["category"] = s.category
	params["side"] = s.side
	params["orderType"] = s.orderType
	params["qty"] = s.size
	if s.price != "" {
		params["price"] = s.price
	}
	if s.orderLinkId != "" {
		params["orderLinkId"] = s.orderLinkId
	}
	if s.positionIdx != nil {
		params["positionIdx"] = *s.positionIdx
	}
	if s.timeInForce != "" {
		params["timeInForce"] = s.timeInForce
	}
	s.setBody(params)

	resp := new(CreateOrderResponse)
	_, err := s.doAuthRequest(ctx, "POST", "/v5/order/create", &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type QueryOrderService struct {
	Service
	pair     string
	category string

	baseCoin   string
	settleCoin string

	orderId     string
	orderLinkId string
}

func (s *QueryOrderService) Pair(c, p string) *QueryOrderService {
	s.pair = p
	s.category = c
	return s
}

func (s *QueryOrderService) OrderID(v string) *QueryOrderService {
	s.orderId = v
	return s
}

func (s *QueryOrderService) ClientOrderID(v string) *QueryOrderService {
	s.orderLinkId = v
	return s
}

func (s *QueryOrderService) BaseCoin(v string) *QueryOrderService {
	s.baseCoin = v
	return s
}

func (s *QueryOrderService) SettleCoin(v string) *QueryOrderService {
	s.settleCoin = v
	return s
}

type QueryOrderResponse struct {
	List           []Order `json:"list"`
	Category       string  `json:"category"`
	NextPageCursor string  `json:"nextPageCursor"`
}

type Order struct {
	OrderId            string  `json:"orderId"`
	OrderLinkId        string  `json:"orderLinkId"`
	BlockTradeId       string  `json:"blockTradeId"`
	Symbol             string  `json:"symbol"`
	Price              float64 `json:"price,string"`
	Qty                float64 `json:"qty,string"`
	Side               string  `json:"side"`
	IsLeverage         string  `json:"isLeverage"`
	PositionIdx        int     `json:"positionIdx"`
	OrderStatus        string  `json:"orderStatus"`
	CancelType         string  `json:"cancelType"`
	RejectReason       string  `json:"rejectReason"`
	AvgPrice           float64 `json:"avgPrice,string"`
	LeavesQty          string  `json:"leavesQty"`
	LeavesValue        string  `json:"leavesValue"`
	CumExecQty         string  `json:"cumExecQty"`
	CumExecValue       string  `json:"cumExecValue"`
	CumExecFee         string  `json:"cumExecFee"`
	TimeInForce        string  `json:"timeInForce"`
	OrderType          string  `json:"orderType"`
	StopOrderType      string  `json:"stopOrderType"`
	OrderIv            string  `json:"orderIv"`
	TriggerPrice       string  `json:"triggerPrice"`
	TakeProfit         string  `json:"takeProfit"`
	StopLoss           string  `json:"stopLoss"`
	TpTriggerBy        string  `json:"tpTriggerBy"`
	SlTriggerBy        string  `json:"slTriggerBy"`
	TriggerDirection   int     `json:"triggerDirection"`
	TriggerBy          string  `json:"triggerBy"`
	LastPriceOnCreated string  `json:"lastPriceOnCreated"`
	ReduceOnly         bool    `json:"reduceOnly"`
	CloseOnTrigger     bool    `json:"closeOnTrigger"`
	SmpType            string  `json:"smpType"`
	SmpGroup           int     `json:"smpGroup"`
	SmpOrderId         string  `json:"smpOrderId"`
	TpslMode           string  `json:"tpslMode"`
	TpLimitPrice       string  `json:"tpLimitPrice"`
	SlLimitPrice       string  `json:"slLimitPrice"`
	PlaceType          string  `json:"placeType"`
	CreatedTime        int64   `json:"createdTime,string"`
	UpdatedTime        int64   `json:"updatedTime,string"`
}

func (s *QueryOrderService) Do(ctx context.Context) (*QueryOrderResponse, error) {
	params := url.Values{}
	params.Add("category", s.category)
	if s.pair != "" {
		params.Add("symbol", s.pair)
	}
	if s.orderLinkId != "" {
		params.Add("orderLinkId", s.orderLinkId)
	}
	if s.orderId != "" {
		params.Add("orderId", s.orderId)
	}
	if s.baseCoin != "" {
		params.Add("baseCoin", s.baseCoin)
	}
	if s.settleCoin != "" {
		params.Add("settleCoin", s.settleCoin)
	}
	s.setQuery(params)
	resp := new(QueryOrderResponse)
	// /v5/order/history 历史订单
	_, err := s.doAuthRequest(ctx, "GET", "/v5/order/realtime", &resp)
	// _, err := s.doAuthRequest(ctx, "GET", "/v5/order/history", &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *Order) Order(category string) *common.Order {
	res := new(common.Order)
	res.ID = r.OrderId
	res.ClientOId = r.OrderLinkId

	// category =
	res.Symbol = r.Symbol
	res.Price = r.Price
	res.Amount = r.Qty
	res.Direction = common.Sell
	if r.Side == "Buy" {
		res.Direction = common.Buy
	}

	switch r.OrderType {
	case "Market":
		res.Type = common.OrderTypeMarket
	case "Limit":
		res.Type = common.OrderTypeLimit
	}
	// if r.PositionIdx > 0 {
	// 	switch r.PositionIdx {
	// 	case 1:
	// 		res.PosSide = exchange.PosSideLong
	// 	case 2:
	// 		res.PosSide = exchange.PosSideLong
	// 	}
	// }

	res.UpdateTime = time.UnixMilli(r.UpdatedTime)
	// res.Cre = time.UnixMilli(r.UpdatedTime)
	switch r.OrderStatus {
	case "New", "Created":
		res.Status = common.OrderStatusNew
	case "Filled", "PartiallyFilledCanceled":
		res.Status = common.OrderStatusFilled
		qty, _ := strconv.ParseFloat(r.CumExecQty, 64)
		res.AvgPrice = r.AvgPrice
		// 卖出时表示实际到账的金额，买入时表示实际花费的金额
		quote := qty * res.AvgPrice

		var fee float64
		if r.CumExecFee == "" {
			var denominator float64
			if res.Direction == common.Buy && category == Spot {
				denominator = qty
			} else {
				denominator = quote
			}
			fee = denominator / 1000
		} else {
			fee, _ = strconv.ParseFloat(r.CumExecFee, 64)
		}

		// 扣除手续费
		if res.Direction == common.Buy {
			if category == Spot {
				qty = qty - fee
			}
		}
		// if res.Direction == common.Buy {
		// 	if category == Spot {
		// 		qty = qty - fee
		// 	} else {
		// 		quote = quote + fee
		// 	}
		// } else {
		// 	quote = quote - fee
		// }
		// res.Quote = quote.RoundFloor(8)
		// res.Result.CommissCoin = r.FeeCcy
		res.Commission = fee
		res.FilledAmount = qty
	case "Cancelled", "Rejected":
		res.Status = common.OrderStatusCancelled
	default:
		res.Status = common.OrderStatusPartiallyFilled
	}
	return res
}
