package lock

import (
	"CoolGoPkg/redis-go/client"
	"gopkg.in/redis.v5"
)

// Locker 锁对象
type Locker struct {
	Rc *redis.Client
}

// NewLocker 新建锁对象
func NewLocker(rc *redis.Client) *Locker {
	lc := &Locker{
		Rc: client.LocalRedis,
	}
	if rc != nil {
		lc.Rc = rc
	}
	return lc
}

// Acquire 获取锁
func (l *Locker) Acquire(key string) bool {
	defaultVal := "lock"
	res, err := l.Rc.SetNX(key, defaultVal, 0).Result()
	if err != nil || !res {
		return false
	}

	return true
}

// Release 释放锁
func (l *Locker) Release(key string) error {
	_, err := l.Rc.Del(key).Result()
	if err != nil {
		return err
	}

	return nil
}
