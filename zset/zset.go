package zset

import (
	"CoolGoPkg/redis-go/client"
	"gopkg.in/redis.v5"
)

const (
	PeopleAndScoreZSetKey = "people_score"
)

var PeopleAndScore = []struct {
	Name  string
	Score float64
}{
	{"xiaoming", 100},
	{"xiaohong", 120},
	{"xiaoyang", 150},
	{"jerry", 60},
}

// 做一个人名，考试分数的例子

// ZSetPeopleAndScore doc
func ZSetPeopleAndScore(members []redis.Z) error {
	return client.LocalRedis.ZAdd(PeopleAndScoreZSetKey, members...).Err()
}

// ZRangePeopleAndScore  分数从小到大 升序
func ZRangePeopleAndScore(start, stop int64) ([]string, error) {
	return client.LocalRedis.ZRange(PeopleAndScoreZSetKey, start, stop).Result()
}

// ZRevRangePeopleAndScore  分数从大到小 降序
func ZRevRangePeopleAndScore(start, stop int64) ([]string, error) {
	return client.LocalRedis.ZRevRange(PeopleAndScoreZSetKey, start, stop).Result()
}

// ZRemPeopleAndScore doc
func ZRemPeopleAndScore(member string) error {
	return client.LocalRedis.ZRem(PeopleAndScoreZSetKey, member).Err()
}
