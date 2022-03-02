package pipeline

import (
	"CoolGoPkg/redis-go/client"
	"fmt"
	"strings"
)

const (
	StringKey1 = "pipeline_string_key1"
	StringKey2 = "pipeline_string_key1"
	ListKey1   = "pipeline_list_key1"
	ListKey2   = "pipeline_list_key2"
	SetKey1    = "pipeline_set_key1"
	SetKey2    = "pipeline_set_key2"
	HashKey1   = "pipeline_hash_key1"
)

func GenBasicDataOfPipeline() {
	err := client.LocalRedis.Set(StringKey1, "stringKey1", 0).Err()
	if err != nil {
		fmt.Println("err : ;", err)
		return
	}
	client.LocalRedis.Set(StringKey2, "stringKey2", 0)
	client.LocalRedis.RPush(ListKey1, 1, 2, 3, 4, 5)
	client.LocalRedis.SAdd(SetKey1, "a", "b", "c")
	client.LocalRedis.HMSet(HashKey1, map[string]string{
		"mark": "1",
		"tom":  "2",
	})
}

func GetDataByPipeline() {
	res, err := client.LocalRedis.Get(StringKey2).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	pipe := client.LocalRedis.Pipeline()
	result1, err := pipe.Get(StringKey1).Result()
	if err != nil {
		fmt.Println("err1 : ", err)
		return
	}
	fmt.Println("result1 : ", result1)

	result2, err := pipe.Get(StringKey2).Result()
	if err != nil {
		fmt.Println("err2 : ", err)
		return
	}
	fmt.Println("result2 : ", result2)

	result3, err := pipe.LRange(ListKey1, 0, -1).Result()
	if err != nil {
		fmt.Println("err3 : ", err)
		return
	}
	fmt.Println("result3 : ", result3)

	result4, err := pipe.SMembers(SetKey1).Result()
	if err != nil {
		fmt.Println("err4 : ", err)
		return
	}
	fmt.Println("result4 : ", result4)

	result5, err := pipe.HGetAll("aaa").Result()
	if err != nil {
		fmt.Println("err5 : ", err)
		return
	}
	fmt.Println("result5 : ", result5)
	resultLast, err := pipe.Exec()
	if err != nil {
		fmt.Println("errlast : ", err)
		return
	}
	fmt.Println(resultLast[4].Err())
	fmt.Println(interface{}(strings.Split(resultLast[3].String(), ":")[1]).([]string)[0])
}
