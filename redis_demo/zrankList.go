package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type Zranking struct {
	Redis          *redis.Client
	Key            string
	Expiration     time.Duration
	StartTimestamp int64
	EndTimestamp   int64
	TimePadwidth   int
}

// rank user
type RankUser struct {
	UID  int64
	Val  int64
	Rank int64
}

func New(rds *redis.Client, key string, exp time.Duration, start int64, end int64) (*Zranking, error) {
	deltaTs := end - start
	if deltaTs < 0 {
		return nil, errors.New("invaild time")
	}
	tpw := len(fmt.Sprint(deltaTs))
	return &Zranking{
		Redis:          rds,
		Key:            key,
		Expiration:     exp,
		StartTimestamp: start,
		EndTimestamp:   end,
		TimePadwidth:   tpw,
	}, nil

}

// Update
func (r *Zranking) Update(ctx context.Context, uid, val int64) (float64, error) {
	valScore, err := r.val2score(ctx, val)
	if err != nil {
		//err = errors.Unwrap(err, "ZRanking Update val2score error")
		return 0, err
	}

	keys := []string{r.Key}
	args := []interface{}{uid, valScore}
	zincrby := redis.NewScript(`
-- 排行榜key
local key = KEYS[1]
-- 要更新的用户id
local uid = ARGV[1]
-- 用户本次新增的val（小数位为时间差标识）
local valScore = ARGV[2]

-- 获取用户之前的score
local score = redis.call("ZSCORE", key, uid)
if score == false then
    score = 0
end
-- 从score中抹除用于时间差标识的小数部分，获取整数的排序val
local val = math.floor(score)

-- 更新用户最新的score信息（累计val.最新时间差）
local newScore = valScore+val
redis.call("ZADD", key, newScore, uid)

-- 更新成功返回newScore（注意要使用tostring才能返回小数）
return tostring(newScore)
	`)
	newScore, err := zincrby.Run(ctx, r.Redis, keys, args...).Float64()
	if err != nil {
		//err = errors.Wrap(err, "ZRanking Update Run lua error")
		return 0, err
	}

	if err := r.Redis.Expire(ctx, r.Key, r.Expiration).Err(); err != nil {
		log.Println("ZRanking Update Expire error. ", "key: ", r.Key, "error: ", err)
	}
	return newScore, nil
}

// 从 score 中获取 val
func (r *Zranking) score2val(ctx context.Context, score float64) (int64, error) {
	scoreStr := fmt.Sprint(score)
	ss := strings.Split(scoreStr, ".")
	valStr := ss[0]
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		//err = errors.Wrap(err, "ZRanking score2val ParseInt error")
		return 0, err
	}
	return val, nil
}

// GetRankingList 返回排行榜
// topN <= 0 取全量
// desc 是否按score降序排列
func (r *Zranking) GetRankingList(ctx context.Context, topN int64, desc bool) ([]RankUser, error) {
	start := int64(0)
	stop := topN - 1
	if topN <= 0 {
		stop = -1
	}

	total := r.GetTotalCount(ctx)
	if stop >= total {
		stop = total - 1
	}

	zrange := r.Redis.ZRangeWithScores
	if desc {
		zrange = r.Redis.ZRevRangeWithScores
	}
	list, err := zrange(ctx, r.Key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	result := []RankUser{}
	for idx, z := range list {
		val, err := r.score2val(ctx, z.Score)
		if err != nil {
			return nil, err
			//return nil, errors.Wrapf(err, "ZRanking GetRankingList score2val error, uid:%v score:%v", z.Member, z.Score)
		}
		member := z.Member.(string)
		uid, err := strconv.ParseInt(member, 10, 64)
		if err != nil {

			return nil, err
			// return nil, errors.Wrapf(err, "ZRanking GetRankingList uid ParseInt error, uid:%v", z.Member)
		}
		m := RankUser{
			UID:  uid,
			Val:  val,
			Rank: int64(idx + 1),
		}
		result = append(result, m)
	}
	return result, nil
}

// GetUserRank 获取某个用户的排行
func (r *Zranking) GetUserRank(ctx context.Context, uid int64, desc bool) (int64, error) {
	zrank := r.Redis.ZRank
	if desc {
		zrank = r.Redis.ZRevRank
	}
	idx, err := zrank(ctx, r.Key, fmt.Sprint(uid)).Result()
	if errors.Is(err, redis.Nil) {
		return 0, nil
	}
	rank := idx + 1
	return rank, err
}

// GetUserVal 获取某个用户score中的排序值
func (r *Zranking) GetUserVal(ctx context.Context, uid int64) (int64, error) {
	score, err := r.Redis.ZScore(ctx, r.Key, fmt.Sprint(uid)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, err
	}
	return r.score2val(ctx, score)
}

// GetTotalCount 获取排行榜总人数
func (r *Zranking) GetTotalCount(ctx context.Context) int64 {
	return r.Redis.ZCard(ctx, r.Key).Val()
}

// val 转为 score:
// score = float64(val.deltaTs)
func (r *Zranking) val2score(ctx context.Context, val int64) (float64, error) {
	nowts := time.Now().Unix()
	deltaTs := r.EndTimestamp - nowts
	scoreFormat := fmt.Sprintf("%%v.%%0%dd", r.TimePadwidth)
	scoreStr := fmt.Sprintf(scoreFormat, val, deltaTs)
	score, err := strconv.ParseFloat(scoreStr, 64)
	if err != nil {
		//err = errors.Wrap(err, "ZRanking val2score ParseFloat error")
		return 0, err
	}
	return score, nil
}
