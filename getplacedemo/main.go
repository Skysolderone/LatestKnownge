package main

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type PairConfig struct {
	ID                 uint    `gorm:"primaryKey;column:id;unsigned" json:"id"`
	Symbol             string  `gorm:"size:32;column:symbol;not null;uniqueIndex:symbol" json:"symbol"`         // 交易对名字
	Platform           string  `gorm:"size:32;column:platform;uniqueIndex:symbol" json:"platform"`              // 平台名（binance,ok,huobi）
	Accuracy           int8    `gorm:"type:int(4) ;column:accuracy;default:0" json:"accuracy"`                  // 数量精度
	Val                float64 `gorm:"type:decimal(32,8) ;column:val;default:0"  json:"val"`                    // 合约面值，仅适用于交割/永续/期权
	SpotMinNotional    float64 `json:"spotMinNotional"`                                                         // 最小现货名义值
	FuturesMinNotional float64 `json:"futuresMinNotional"`                                                      // 最小合约名义值
	AmountAccuracy     int8    `gorm:"type:int(4) ;column:amount_accuracy;default:0" json:"amount_accuracy"`    // 下单量精度
	ETF                bool    `gorm:"type:int(4) unsigned;column:etf;default:0" json:"etf"`                    // ETF资产
	Trace              bool    `gorm:"type:int(4) unsigned;column:trace;default:0" json:"trace"`                // 是否支持带单
	Future             uint8   `gorm:"type:int(4) unsigned;column:future;default:0" json:"future"`              // 0:只支持现货，1:支持现货和合约，2：只支持合约，3：热门交易对，趋同
	Leverage           uint8   `gorm:"type:int(4) unsigned;column:leverage;default:0" json:"leverage"`          // 最大杠杆倍数
	Hot                bool    `gorm:"type:int(4) unsigned;column:hot;default:0" json:"hot"`                    // 热门
	Status             uint8   `gorm:"type:int(4) unsigned;column:status;index:status;default:0" json:"status"` // 0，正常，1 交易被下架

	Category uint8 `gorm:"default:4" json:"category"` // 交易对分类：1(杠杆：4x、5x、6x、7x、8x) 2(杠杆：2x、3x、4x、5x、6x) 3(杠杆: 1x、2x、3x、4x、5x)

	Market      bool    `gorm:"type:int(4) unsigned;column:market;default:0" json:"market"` // 是否显示在行情页, 是否开启手动交易
	Autrade     uint8   `gorm:"default:7" json:"autrade"`                                   // 自动交易: 0x01 = 现货, 0x02 = 多, 0x04 = 空
	Sigtrade    uint8   `gorm:"default:7" json:"sigtrade"`                                  // 信号交易: 0x01 = 现货, 0x02 = 多, 0x04 = 空
	AdminSignal uint8   `gorm:"default:1" json:"admin_signal"`                              // Admin 信号: 0x01 = 现货
	Cover       uint8   `gorm:"default:7" json:"cover"`
	CoverRatio  float64 `gorm:"default:1" json:"cover_ratio"`
	Strategy    string  `gorm:"size:64;column:strategy;default:1,5,8,9,13,15,17,18" json:"strategy"` // 支持的特殊策略ID，以,隔开
	ListingTime int64   `gorm:"default:0" json:"listing_time"`                                       // 上架时间

	// 0.其他; 1.Infrastructure   公链v2.DeFi   去中心化金融; 3.Storage 存储; 4.POS  权益证明; 5.POW 工作量证明
	// 6. Layer 1/Layer 2  （一层/二层网络）; 7.Metaverse 元宇宙; 8.Gaming  游戏; 9.NFT  非同质化代币; 10.Fan Token 粉丝代币;
	// 11.Polkadot  波卡生态; 12.BNB Chain 币安链; 13.Innovation  创新区; 14.Launchpad 新币质押; 15.Launchpoo  新币挖矿
	Categories string `gorm:"size:32;column:categories;default:''" json:"categories"` // 分类，以,隔开
	Time       int64  `json:"time"`
}

func main() {
	dsn := "root:gg123456@tcp(127.0.0.1:3306)/trading?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	resp := PairConfig{}
	db.Table("pair_configs").Where("id=1461").First(&resp)
	fmt.Println(resp.Categories)
	plates := make([]int, 0)
	if resp.Categories == "" {
		fmt.Println(plates)
	}
	s := strings.ReplaceAll(resp.Categories, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	cate := strings.Split(s, ",")
	for _, v := range cate {
		fmt.Println(v)
		s, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		plates = append(plates, s)
	}
	fmt.Println(plates)
}
