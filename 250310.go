package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

func LocalModel[T comparable, key any, value comparable](T) {
}

type user struct {
	Name string
	Age  int
}
type book struct {
	Name string
	Age  int
}
type ModelType interface {
	user | book
}
type KeyModel interface {
	int | float64
}

var testm sync.Map

func load_model[T ModelType, key KeyModel](model T, k key, m *sync.Map) (T, error) {
	v, ok := m.Load(k)
	if !ok {
		return model, errors.New("not found")
	}
	return v.(T), nil
}

func main() {
	s1 := user{
		Name: "testuser",
		Age:  18,
	}
	s2 := book{
		Name: "testbook",
		Age:  20,
	}
	testm.Store(1, s1)
	// testm.Load(1)
	// testm.Delete(1)
	testm.Store(2, s2)
	t1, err := load_model(user{}, 1, &testm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1)
	t2, err := load_model(book{}, 2, &testm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)

	fmt.Println(math.Abs(-1.0))
	now := time.Now()
	fmt.Println(now)
	ts := now.Truncate(time.Hour * 8).Add(time.Hour * 8)
	fmt.Println(time.Until(ts))

	fmt.Println(timeUntilNext(timperiod1))
	fmt.Println(timeUntilNext(timeperiod2))
}

var (
	timperiod1  = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24}
	timeperiod2 = []int{4, 8, 12, 16, 20, 24}
)

// 计算当前时间到下一个目标时间的剩余时间
func timeUntilNext(periods []int) time.Duration {
	now := time.Now()
	currentHour := now.Hour()
	loc := now.Location()

	for _, hour := range periods {
		if currentHour < hour { // 找到下一个最近的时间点
			targetTime := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, loc)
			return time.Until(targetTime)
		}
	}
	// 如果今天已经过了所有时间点，计算明天的第一个时间点
	targetTime := time.Date(now.Year(), now.Month(), now.Day()+1, periods[0], 0, 0, 0, loc)
	return time.Until(targetTime)
}
