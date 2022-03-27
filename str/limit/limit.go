package limit

import (
	"gopkg.in/redis.v5"
	"strconv"
	"time"
)

// Limiter 限速器
type Limiter struct {
	RClient    *redis.Client
	Key        string
	LimitTimes int
	Expr       time.Duration
}

// NewLimiter 新建一个限速器
func NewLimiter(rc *redis.Client, key string, times int, exp time.Duration) *Limiter {
	return &Limiter{
		RClient:    rc,
		Key:        key,
		LimitTimes: times,
		Expr:       exp,
	}
}

// SetMaxExecuteTimes doc
func (l *Limiter) SetMaxExecuteTimes(times int) error {
	ts := l.LimitTimes
	if times > 0 {
		ts = times
	}
	return l.RClient.Set(l.Key, ts, l.Expr).Err()
}

// StillValidToExecute 是否仍可执行
func (l *Limiter) StillValidToExecute() bool {
	num, err := l.RClient.Decr(l.Key).Result()
	if err != nil {
		return false
	}

	return num >= 0
}

// GetRemainExecuteTimes 获取剩余的执行次数
func (l *Limiter) GetRemainExecuteTimes() (int, error) {
	numStr, err := l.RClient.Get(l.Key).Result()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(numStr)
}
