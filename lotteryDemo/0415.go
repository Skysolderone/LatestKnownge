package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Prize 定义奖品类型
type Prize struct {
	Name   string // 奖品名称
	Number int    // 奖品数量
}

// 初始化奖品设置
var prizes = []Prize{
	{"特等奖：iPhone18 1TB", 1},
	{"一等奖：华为笔记本电脑", 2},
	{"二等奖：索尼相机", 3},
	{"三等奖：小米烤箱", 5},
	{"四等奖：小米加湿器", 8},
	{"五等奖：枕头", 10},
	{"六等奖：零食大礼包", 20},
}

// 抽奖概率
type Probability struct {
	RechargeAbility float64 // 用户充值能力大于5000的概率
	ActivityLevel   float64 // 用户每日活跃度概率
	WealthLevel     float64 // 用户总财富等级大于28级的概率
}

// 初始化默认概率
var defaultProbability = Probability{
	RechargeAbility: 0.2,
	ActivityLevel:   0.3,
	WealthLevel:     0.5,
}

// 用户信息
type User struct {
	ID          int     // 用户ID
	Recharge    float64 // 本月充值金额
	Activity    float64 // 日活跃时间
	WealthLevel int     // 财富等级
}

// 检查用户本月充值能力是否大于5000
func isUserCharge(userID int) bool {
	// 伪代码
	// return database.UserCharge(userID) > 5000
	return false
}

// 检查用户每日活跃度是否大约5个小时
func isUserActive(userID int) bool {
	// 伪代码
	// return database.UserActivity(userID) >= 5
	return false
}

// 检查用户总的财富等级是否大于28级
func isUserWealth(userID int) bool {
	// 伪代码
	// return database.UserWealth(userID) > 28
	return false
}

// 检查用户是否中奖
func checkWinning(user User, probability Probability) bool {
	// 计算用户是否中特等奖
	if user.Recharge > 5000 && rand.Float64() < probability.RechargeAbility {
		return true
	}
	// 计算用户是否中其他奖项
	if user.Activity >= 5 && rand.Float64() < probability.ActivityLevel {
		return true
	}
	if user.WealthLevel > 28 && rand.Float64() < probability.WealthLevel {
		return true
	}
	return false
}

// 抽奖函数
func drawPrize(user User, probability Probability) string {
	// 检查用户是否中奖
	if checkWinning(user, probability) {
		// 随机选择一个奖品
		prizeIndex := rand.Intn(len(prizes))
		// 更新奖品数量
		if prizes[prizeIndex].Number > 0 {
			prizes[prizeIndex].Number--
			// 返回中奖信息
			return fmt.Sprintf("恭喜用户%d中奖，获得%s", user.ID, prizes[prizeIndex].Name)
		} else {
			// 如果奖品已经抽完，则返回未中奖信息
			return fmt.Sprintf("用户%d未中奖", user.ID)
		}
	}
	// 未中奖
	return fmt.Sprintf("用户%d未中奖", user.ID)
}

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 模拟用户数据
	users := []User{
		{1, 6000, 6, 30},
		{2, 2000, 3, 25},
		{3, 4000, 8, 35},
	}

	// 启动并发抽奖
	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go func(user User) {
			defer wg.Done()
			// 每个用户抽取100次奖品
			for i := 0; i < 20; i++ {
				// 模拟用户抽奖
				fmt.Println(drawPrize(user, defaultProbability))
			}
		}(user)
	}
	wg.Wait()
}
