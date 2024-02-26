package bybit

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
	"trading/utils"
	"trading/utils/websocket"

	"github.com/dolotech/log"
)

type Option func(*Client)

func SetUseOptionsURL() Option {
	return func(c *Client) {
		c.baseURL = optionsBaseURL
		// c.baseWsURL = baseWsAwsURL
	}
}

func SetBaseURL(v string) Option {
	return func(c *Client) {
		c.baseURL = v
	}
}

func SetProxy(proxyURL string) Option {
	return func(c *Client) {
		if proxyURL == "" {
			return
		}
		if !strings.Contains(proxyURL, "http://") && !strings.Contains(proxyURL, "https://") {
			proxyURL = "http://" + proxyURL
		}
		c.baseURL = proxyURL
		c.header.Set("x-ex", "bybit")
	}
}

func SetReferer(referer string) Option {
	return func(c *Client) {
		c.header.Add("referer", referer)
	}
}

type Client struct {
	baseURL   string
	baseWsURL string
	Debug     bool
	header    http.Header
}

func NewClient(opts ...Option) *Client {
	clt := new(Client)
	clt.header = http.Header{}
	for _, opt := range opts {
		opt(clt)
	}
	if clt.baseURL == "" {
		clt.baseURL = baseURL
	}
	if clt.baseWsURL == "" {
		clt.baseWsURL = baseWsURL
	}

	return clt
}

func (c *Client) NewInfoService() *InfoService {
	s := new(InfoService)
	s.c = c
	return s
}

func (c *Client) NewKlineService() *KlineService {
	s := new(KlineService)
	s.c = c
	return s
}

func (c *Client) NewBalanceService() *BalanceService {
	s := new(BalanceService)
	s.c = c
	return s
}

func (c *Client) NewApiService() *ApiService {
	s := new(ApiService)
	s.c = c
	return s
}

func (c *Client) NewUpgradeUTAService() *UpgradeUTAService {
	s := new(UpgradeUTAService)
	s.c = c
	return s
}

func (c *Client) NewConfigService() *ConfigService {
	s := new(ConfigService)
	s.c = c
	return s
}

func (c *Client) NewOrderService() *OrderService {
	s := new(OrderService)
	s.c = c
	return s
}

func (c *Client) NewQueryOrderService() *QueryOrderService {
	s := new(QueryOrderService)
	s.c = c
	return s
}

func (c *Client) NewChangeLeverageService() *ChangeLeverageService {
	s := new(ChangeLeverageService)
	s.c = c
	return s
}

func (c *Client) NewChangeMarginModeService() *ChangeMarginModeService {
	s := new(ChangeMarginModeService)
	s.c = c
	return s
}

func (c *Client) NewPositionService() *PositionService {
	s := new(PositionService)
	s.c = c
	return s
}

func (c *Client) NewChangePositionModeService() *ChangePositionModeService {
	s := new(ChangePositionModeService)
	s.c = c
	return s
}

func (c *Client) NewPublicWsClient(category string) *PublicWsClient {
	pwc := &PublicWsClient{
		c:             c,
		topicHandlers: make(map[string]TopicHandlerFunc, 1024),
		operationHandlers: map[string]func(*OperationResp){
			"subscribe":   func(ds *OperationResp) { log.Info("subscribe", "arg", ds.Arg) },
			"unsubscribe": func(ds *OperationResp) { log.Info("unsubscribe", "args", ds.Arg) },
		},
	}
	endpoint := c.baseWsURL
	switch category {
	case Spot:
		endpoint += spotWsPath
	case Futures:
		endpoint += futuresWsParh
	}
	pwc.wc = websocket.NewClient(pwc, endpoint)
	pwc.register("ping", func(ds *OperationResp) {
		pwc.lastPong = time.Now()
	})
	return pwc
}

const pongTimeout = 30 * time.Second

var (
	pingText = []byte(`{"op": "ping"}`)
)

type TopicHandlerFunc func(data []byte)

type TopicFrame struct {
	Topic string          `json:"topic"`
	Type  string          `json:"type"`
	Data  json.RawMessage `json:"data"`
	Ts    int64           `json:"ts"`
	Cs    int64           `json:"cs"`
}

type OperationReq struct {
	Op   string `json:"op"` //
	Args []any  `json:"args"`
}

type OperationResp struct {
	WsErr
	Op  string          `json:"op"`
	Arg json.RawMessage `json:"arg"`
}

type PublicWsClient struct {
	c  *Client
	wc *websocket.Client

	topicHandlers map[string]TopicHandlerFunc
	lastPong      time.Time

	operationHandlers map[string]func(*OperationResp)
}

func (c *PublicWsClient) OnMessage(conn *websocket.Conn, data []byte) {
	err := c.handleMessage(data)
	if err != nil {
		log.Error("handle message", err, "data", string(data))
	}
}

func (c *PublicWsClient) OnErr(err error) {
	log.Error("websocket read", err)
}

func (c *PublicWsClient) MessageType() int {
	return websocket.TextMessage
}

func (c *PublicWsClient) Start() {
	for {
		err := c.wc.Dial()
		if err != nil {
			log.Error("dial", err)
			continue
		}
		c.flush() // 发出所有的订阅消息
		c.lastPong = time.Now()
		ticker := time.NewTicker(time.Duration((20 + utils.RandInt(8))) * time.Second)
		for {
			select {
			case <-c.wc.Done():
				continue
			case <-c.wc.Stopped():
				return
			case <-ticker.C:
				if time.Since(c.lastPong) > pongTimeout {
					log.Error("server respone timeout")
				}
				if err := c.wc.Send(pingText); err != nil {
					log.Error("ping", err)
				}
			}
		}
	}
}

func (c *PublicWsClient) Stop() {
	c.wc.Stop()
}

func (c *PublicWsClient) handleMessage(data []byte) error {
	tf := new(TopicFrame)
	_ = json.Unmarshal(data, tf)
	if tf.Topic != "" {
		return c.handleTopic(tf)
	}

	op := new(OperationResp)
	err := json.Unmarshal(data, op)
	if err != nil {
		return c.handleOperation(op)
	}
	return err
}
func (c *PublicWsClient) handleOperation(op *OperationResp) error {
	if op.Op != "" {
		return errors.New("missing op of message")
	}
	handler, ok := c.operationHandlers[op.Op]
	if !ok {
		return errors.New("unhandled operation")
	}
	handler(op)
	return nil
}

func (c *PublicWsClient) handleTopic(msg *TopicFrame) error {
	handler, ok := c.topicHandlers[msg.Topic]
	if !ok {
		return errors.New("unhandled topic")
	}
	handler(msg.Data)
	return nil

}

func (c *PublicWsClient) flush() {
	msg := OperationReq{
		Op:   "subscribe",
		Args: make([]any, 0, len(c.topicHandlers)),
	}
	for key := range c.topicHandlers {
		msg.Args = append(msg.Args, key)
	}
	err := c.wc.JSON(msg)
	if err != nil {
		log.Error("subscribe", err)
	}
}

func (c *PublicWsClient) register(op string, handler func(*OperationResp)) {
	c.operationHandlers[op] = handler
}

func (c *PublicWsClient) sub(topic string, handler TopicHandlerFunc) error {
	c.topicHandlers[topic] = handler
	return nil
}
