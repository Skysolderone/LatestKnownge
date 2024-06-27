package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Player struct {
	ID            uint    `gorm:"primaryKey;column:id;unsigned"`                                             // 用户唯一ID
	Phone         string  `gorm:"size:32;column:phone;default:''"`                                           // 手机号
	Mail          string  `gorm:"size:64;column:mail;default:'';uniqueIndex:index_mail"`                     // 邮箱地址
	Nickname      string  `gorm:"size:64;column:nickname;default:''"`                                        // 昵称
	Wallet        string  `gorm:"size:64;column:wallet;default:'';index:index_wallet"`                       // TRC20-USDT钱包地址
	PolygonWallet string  `gorm:"size:64;column:polygon_wallet;default:'';index:index_polygon_wallet"`       // Polygon钱包地址
	Balance       float64 `gorm:"type:decimal(18,3) unsigned;column:balance;default:0"`                      // TRC20-USDT钱包余额
	Energy        float64 `gorm:"type:decimal(18,5) ;column:energy;default:0"`                               // 小蚁能量
	Password      string  `gorm:"size:64;column:password;default:''"`                                        // 登录密码
	PayPassword   string  `gorm:"size:64;column:pay_password;default:''"`                                    // 支付密码
	InviteID      uint    `gorm:"type:int(11) unsigned;column:invite_id;default:0"`                          // 邀请人的ID
	OwnInvite     string  `gorm:"size:6;column:own_invite;default:'';not null;uniqueIndex:index_own_invite"` // 玩家自身的邀请码
	Photo         string  `gorm:"size:1024;column:photo;default:''"`                                         // 头像
	Vip           uint8   `gorm:"type:int(4) unsigned;column:vip;default:0"`                                 // VIP等级
	TVip          uint8   `gorm:"type:int(4) unsigned;column:tvip;default:0"`                                // TVIP等级
	LoginTime     uint    `gorm:"type:int(11) unsigned;column:login_time;default:0"`                         // 最后一次登录时间戳
	RegistTime    uint    `gorm:"type:int(11) unsigned;column:regist_time;default:0"`                        // 注册时间戳
	RegistIP      uint    `gorm:"type:int(11) unsigned;column:regist_ip;default:0"`                          // 注册IP地址
	LoginIP       uint    `gorm:"type:int(11) unsigned;column:login_ip;default:0"`                           // 最后一次登录IP地址
	ServeID       uint8   `gorm:"type:int(4);column:serve_id;default:0"`                                     // 服务器ID
	Remark        string  `gorm:"size:128;column:remark;default:''"`                                         // 备注
	Addr          string  `gorm:"size:128;column:addr;default:''"`                                           // 地址
	Agency        uint8   `gorm:"type:int(4) unsigned;column:agency;default:0"`                              // 0:普通用户，1:代理
	Commision     float64 `gorm:"type:decimal(18,3) unsigned;column:commision;default:0"`                    // 佣金比例
	Device        string  `gorm:"size:128;column:device;default:''"`                                         // 设备ID

	Share          float64 `gorm:"type:decimal(18,4) unsigned;column:share;default:0.5"`        // 总投入控制
	ShareFuture    float64 `gorm:"type:decimal(18,4) unsigned;column:share_future;default:0.5"` // 期货开仓仓位水平（0<N<=1 默认0.25  0也就是0.25）多
	ShareSpot      float64 `gorm:"type:decimal(18,3) unsigned;column:share_spot;default:0"`     // 现货开仓仓位水平（0<N<=1 默认0.25  0也就是0.25）
	Contect        string  `gorm:"size:128;column:contect;default:''"`                          // 联系方式
	ContectType    string  `gorm:"size:128;column:contect_type;default:''"`                     // 联系方式类型 (Phone,Email,WhastApp,Telegram,Wechat)
	ContectVisible uint8   `gorm:"type:int(4) unsigned;column:contect_visible;default:1"`       // 下级是否可见（0：不可见，1：下级可见）
	SignalExpire   uint    `gorm:"type:int(11) unsigned;column:signal_expire;default:0"`        // 智能信号续期
	// 从energy_logs表中同步
	Activated            uint8 `gorm:"type:int(4) unsigned;column:activated;default:0"`               // 激活（0：未激活，1：已激活）
	ActivateTime         uint  `gorm:"type:int(11) unsigned;column:activate_time;default:0"`          // 激活时间戳
	RegistType           uint8 `gorm:"type:int(4) unsigned;column:regist_type;default:1"`             // 1手机，2落地页，3系统生成, 4 电报注册
	LastUpgradedReferral int   `gorm:"type:int(11) unsigned;column:last_upgraded_referral;default:0"` // 上次升级时的直推用户数，用作vip升级功能 因为升级需要的条件是新邀请用户的，新直推用户=总的直推用户-LastUpgradedReferral
	UpgradeTime          uint  `gorm:"default:0"`                                                     // 升级时间

	LockTime  uint  // 用户验证锁定时间
	DeletedAt int64 // 删除时间
	Status    uint8 `gorm:"default:0"` // 0 正常，1 注销中（7天），2 已注销

	ChangeInviter uint8 `gorm:"type:int(4) unsigned;column:change_inviter;default:0"` // 改变上级
	// 智能仓位管理、订阅等状态，可提供64种类型
	// 1 << 0 智能仓位，Substatus二进制值右一 0关闭，1开启
	Substatus       uint64 `gorm:"type:bigint(64) unsigned;column:substatus;default:0"`
	SuperiorVisible uint8  `gorm:"column:superior_visible;default:7"` // 上级是否可见（0：不可见，1：资金可见，2: 机器人可见,4: 联系方式可见）

	GuidanceArea string `gorm:"size:8;default:''"` // 指导地区码，设置了这个字段就可以在APP的‘寻找教练’中列出
	/*
		指导语言:
			英语 = "en",
			法语 = "fr",
			印尼语 = "id",
			印地语 = "hi",
			俄语 = "ru",
			韩语 = "ko",
			西班牙语 = "es",
			葡萄牙语 = "pt",
			阿拉伯语 = "ar",
			德语 = "de",
			越南语 = "vi",
			日语 = "ja",
			意大利语 = "it",
	*/
	GuidanceLang string `gorm:"size:8;default:''"`

	Authenticator string `gorm:"size:256;default:''"` // 验证器密码

	PendingEnergy float64 `gorm:"type:decimal(18,5) ;column:pending_energy;default:0"` // 待领取能量
	JoinComission uint8   `gorm:"default:1"`                                           // 启用收益排行 1 enabled 0 unenabled
	JoinInvite    uint8   `gorm:"default:1"`                                           // 启用直推排行
	JoinIndirect  uint8   `gorm:"default:1"`                                           // 启用间推排行

	activated         *bool  `gorm:"-"`
	activateCountdown int64  `gorm:"-"`
	Country           string `gorm:"size:256;default:''" json:"country"`

	ContactMail string `gorm:"size:64;column:contact_mail;default:''" json:"contact_mail"` // 联系邮箱
	Telegram    string `gorm:"size:256;default:''" json:"telegram"`
	TgGroup     string `gorm:"size:256;default:''" json:"tg_group"` // 群组联系方式
	WaGroup     string `gorm:"size:256;default:''" json:"wa_group"`
	WhatsApp    string `gorm:"size:256;default:''" json:"whats_app"`
	FaceBook    string `gorm:"size:256;default:''" json:"face_book"`
	Twitter     string `gorm:"size:256;default:''" json:"twitter"`
	Youtobe     string `gorm:"size:256;default:''" json:"youtobe"`
	Tiktok      string `gorm:"size:256;default:''" json:"tiktok"`

	Earnest  uint `gorm:"type:int(11) unsigned;column:earnest;default:0"`  // 1  1000 2 2000 3 3000 缴纳保证金等级
	Refunded uint `gorm:"type:int(11) unsigned;column:refunded;default:0"` // default 0 1 申请退款中 2已经退款
}

func main() {
	dsn := "root:gg123456@tcp(172.22.0.1:3306)/trading?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.Debug()
	if err != nil {
		log.Fatal(err)
	}
	p := Player{}
	s := "+12 345678"
	db.Table("players").
		Where(Player{ID: 21, Nickname: "89k"}).
		Or(Player{Phone: s, Nickname: "89"}).
		Or(Player{Phone: s, Nickname: "8l"}).Find(&p)
	fmt.Println(p)
}
