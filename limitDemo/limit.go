package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	limiter_gin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

var globalRate = limiter.Rate{
	Period: 1 * time.Hour,
	Limit:  1000,
}
var routeNameMap = map[string]string{
	"/api/users": "users",
	"/api/items": "items",
}

var rateMapJSON = `{
	"default:users": "1000-M",
	"strict:users": "10-S",
	"default:items": "1000-M",
	"strict:items": "10-S"
   }`

type RateConfig map[string]string

func getRateConfigFromDB(mode, routeName string) (string, error) {
	return rateMapJSON, nil
}

func parseRate(rateStr string) (limiter.Rate, error) {
	parts := strings.Split(rateStr, "-")
	if len(parts) != 2 {
		return limiter.Rate{}, fmt.Errorf("invalid rate format:%s", rateStr)
	}
	limit, err := limiter.NewRateFromFormatted(rateStr)
	if err != nil {
		return limiter.Rate{}, err
	}
	return limit, nil
}
func retrieveRateConfig(mode, routeName string) (limiter.Rate, error) {
	RateControlJSON, err := getRateConfigFromDB(mode, routeName)
	if err != nil {
		return limiter.Rate{}, err
	}
	rateConfig := make(RateConfig)
	err = json.Unmarshal([]byte(RateControlJSON), &rateConfig)
	if err != nil {
		return limiter.Rate{}, err
	}
	rateStr, exists := rateConfig[mode+":"+routeNameMap[routeName]]
	if !exists {
		return limiter.Rate{}, fmt.Errorf("rate configuration not found for mode :%s,routeName:%s", mode, routeName)
	}
	return parseRate(rateStr)
}

func RateControl(c *gin.Context) {
	routeName := c.FullPath()
	mode := "default"
	rate, err := retrieveRateConfig(mode, routeName)
	if err != nil {
		rate = globalRate
	}
	storeWithPrefix := memory.NewStoreWithOptions(
		&memory.Options{
			Prefix:   mode + ":" + routeName + ":",
			MaxRetry: 3,
		},
	)
	rateLimiter := limiter.New(storeWithPrefix, rate)
	limiter_gin.RateLimiter(rateLimiter).Middleware(c)

}
func main() {
	r := gin.Default()
	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Users route"})
	})
	r.Use(RateControl)
	r.GET("/api/items", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Items route"})
	})
	r.Run(":8080")
}
