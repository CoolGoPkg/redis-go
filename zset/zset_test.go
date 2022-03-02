package zset

import (
	"CoolGoPkg/redis-go/client"
	"gopkg.in/redis.v5"
	"testing"
)

//func init() {
//	client.Init()
//}

func TestZSetPeopleAndScore(t *testing.T) {
	client.Init()
	members := make([]redis.Z, len(PeopleAndScore))
	for i, ps := range PeopleAndScore {
		members[i] = redis.Z{
			Member: ps.Name,
			Score:  ps.Score,
		}
	}
	err := ZSetPeopleAndScore(members)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	t.Log("success !")
}

func TestZRangePeopleAndScore(t *testing.T) {
	client.Init()
	member, err := ZRangePeopleAndScore(0, 3)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	t.Log("member", member)
}

func TestZRevRangePeopleAndScore(t *testing.T) {
	client.Init()
	member, err := ZRevRangePeopleAndScore(0, 3)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	t.Log("member", member)
}

func TestZRemPeopleAndScore(t *testing.T) {
	client.Init()
	err := ZRemPeopleAndScore("xiaoming")
	if err != nil {
		t.Log("err : ", err)
		return
	}

	t.Log("success ")
}
