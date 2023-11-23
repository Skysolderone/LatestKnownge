package ratelimit

import "sync"

type Leak struct {
	limit    int //总大小
	rate     int //速率
	current  int //当前位置
	mute     sync.Mutex
	timesamp time.Time
}

func NewLeak(limint, rate int) *Leak {
	return &Leak{limit: limint, rate: rate, timesamp: time.Now()}
}

func (l *Leak) try() bool {
	l.mute.Lock()
	defer l.mute.Unlock()
	now := time.Now()
	interval := time.Sub(l.timesamp)
	if interval > time.Second() {
		l.current = max(0, l.current-int(interval/time.Second())*l.rate)
		l.timesamp = now
	}
	if l.current >= l.limit {
		return false
	}
	t.current++
	return trues
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
