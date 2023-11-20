package ratelimit

import (
	"sync"
	"time"
)

type Token struct {
	limit    int
	rate     int
	current  int
	mute     sync.Mutex
	timesamp time.Time
}

func NewLeak(limit, rate int) *Token {
	return &Token{limit: limit, rate: rate, timesamp: time.Now()}
}

func (t *Token) try() bool {
	t.mute.Lock()
	defer t.mute.Unlock()
	now := time.Now()
	interval := time.Sub(t.timesamp)
	if interval > time.Second() {
		t.current = min(t.limit, t.current+int(interval/time.Second)*t.rate)
		t.timesamp = now
	}
	if t.current == 0 {
		return false
	}
	t.current--
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
