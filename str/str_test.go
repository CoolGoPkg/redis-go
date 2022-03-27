package str

import (
	"CoolGoPkg/redis-go/client"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestSetStr(t *testing.T) {
	client.Init()
	err := SetStr("mark", "123")
	if err != nil {
		t.Log("set err : ", err)
		return
	}

	res, err := GetStr("mark")
	if err != nil {
		t.Log("get err : ", err)
		return
	}

	t.Log("success get result : ", res)
}

func TestSetStrNX(t *testing.T) {
	client.Init()
	err := SetStrNX("name", "mark")
	if err != nil {
		t.Log("set nx 1 err : ", err)
		return
	}

	res, err := GetStr("name")
	if err != nil {
		t.Log("get err : ", err)
		return
	}
	t.Log("success get result : ", res)

	err = SetStrNX("name", "mark")
	if err != nil {
		t.Log("set nx 2 err : ", err)
		return
	}

	t.Log("success get result : ", res)
}

func TestSetStrXX(t *testing.T) {
	client.Init()
	err := SetStr("num", "1")
	if err != nil {
		t.Log("set err : ", err)
		return
	}

	err = SetStrXX("num", "2")
	if err != nil {
		t.Log("set xx 1 err : ", err)
		return
	}

	res, err := GetStr("xxx")
	if err != nil {
		t.Log("get err : ", err)
		return
	}
	t.Log("success get result : ", res)
}

func TestGetStr(t *testing.T) {
	client.Init()
	res, err := GetStr("password")
	if err != nil {
		t.Log("get err : ", err)
		return
	}
	t.Log("success get result : ", res)
}

func TestGetSetStr(t *testing.T) {
	client.Init()
	res, err := GetSetStr("pass", "123456")
	if err != nil {
		t.Log("getset err : ", err)
		return
	}

	t.Log("success getset result : ", res)
}

func TestMSetStr(t *testing.T) {
	client.Init()
	res, err := MSetStr()
	if err != nil {
		t.Log("mset err : ", err)
		return
	}

	key1, err := GetStr("key3")
	if err != nil {
		t.Log("mset err : ", err)
		return
	}

	key2, err := GetStr("key4")
	if err != nil {
		t.Log("mset err : ", err)
		return
	}

	key3, err := GetStr("key5")
	if err != nil {
		t.Log("mset err : ", err)
		return
	}
	key6, err := GetStr("key6")
	if err != nil {
		t.Log("mset err : ", err)
		return
	}

	t.Log("success mset result : ", res, key1, key2, key3, key6)
}

func TestMGetStr(t *testing.T) {
	client.Init()
	res, err := MGetStr()
	if err != nil {
		t.Log("mset err : ", err)
		return
	}
	for _, r := range res {
		fmt.Printf("%T \n", r)
	}

	t.Log("success mget result : ", res)
}

type MySlice []string

func (ms MySlice) MarshalBinary() ([]byte, error) {
	return json.Marshal(ms)
}

func TestSetInterface(t *testing.T) {
	client.Init()
	s := Student{"abc", 123}
	res, err := client.LocalRedis.Set("student", s, 0).Result()
	if err != nil {
		t.Log("redis set err : ", err)
		return
	}

	t.Log("set success result : ", res)

	msetRes, err := client.LocalRedis.MSet("int", 0, "float", 3.14, "myslice", MySlice{"a", "b", "c"}).Result()
	if err != nil {
		t.Log("redis Mset err : ", err)
		return
	}

	t.Log("Mset success result : ", msetRes)

}

func TestGetInterface(t *testing.T) {
	client.Init()
	mgetRes, err := client.LocalRedis.MGet("int", "float", "student", "myslice").Result()
	if err != nil {
		t.Log("redis Mset err : ", err)
		return
	}

	for i, mr := range mgetRes {
		fmt.Printf("%T \n", mr)
		if i == 0 {
			mrInt, err := strconv.Atoi(mr.(string))
			if err != nil {
				t.Log("atoi err :", err)
				return
			}
			fmt.Printf("int : %T,%d \n", mrInt, mrInt)
		}

		if i == 1 {
			mrfloat, err := strconv.ParseFloat(mr.(string), 64)
			if err != nil {
				t.Log("ParseFloat err :", err)
				return
			}
			fmt.Printf("float : %T,%f \n", mrfloat, mrfloat)
		}

		if i == 2 {
			ms := Student{}
			err := json.Unmarshal([]byte(mr.(string)), &ms)
			if err != nil {
				t.Log("json unmarshal student err :", err)
				return
			}
			fmt.Println("student", ms)
		}
		if i == 3 {
			ms := MySlice{}
			err := json.Unmarshal([]byte(mr.(string)), &ms)
			if err != nil {
				t.Log("json unmarshal myslice err :", err)
				return
			}
			fmt.Println("myslice", ms)
		}

	}
}

func TestMSetNXStr(t *testing.T) {
	client.Init()
	ok, err := client.LocalRedis.MSetNX("a", "a", "f", "f").Result()
	if err != nil {
		t.Log("redis MSetNX err : ", err)
		return
	}
	t.Log("set nx ", ok)

	a, err := GetStr("a")
	if err != nil {
		t.Log("mset err : ", err)
		return
	}
	t.Log("a", a)

	f, err := GetStr("f")
	if err != nil {
		t.Log("mset err : ", err)
		return
	}

	t.Log("f", f)

}

func TestGetStrLen(t *testing.T) {
	client.Init()
	key := "strLen_not_find"
	//val := "abc"
	//res, err := client.LocalRedis.Set(key, val, 0).Result()
	//if err != nil {
	//	t.Log("redis set err : ", err)
	//	return
	//}
	//t.Log("set success result : ", res)
	l, err := GetStrLen(key)
	if err != nil {
		t.Log("redis getLen err : ", err)
		return
	}

	t.Log("length :  ", l)

}

func TestGetStrRange(t *testing.T) {
	client.Init()
	key := "strRange"
	val := "0123456789"
	res, err := client.LocalRedis.Set(key, val, 0).Result()
	if err != nil {
		t.Log("redis set err : ", err)
		return
	}
	t.Log("set success result : ", res)
	// 从头部开始
	res, err = GetStrRange(key, 0, 3)
	if err != nil {
		t.Log("redis get range 1 err : ", err)
		return
	}
	t.Log("get range 1 success result : ", res)

	// 头尾混取
	res, err = GetStrRange(key, 1, -2)
	if err != nil {
		t.Log("redis get range 2 err : ", err)
		return
	}
	t.Log("get range 2 success result : ", res)

	// 从尾部开始取
	res, err = GetStrRange(key, -3, -1)
	if err != nil {
		t.Log("redis get range 3 err : ", err)
		return
	}
	t.Log("get range 3 success result : ", res)

	// 取所有
	res, err = GetStrRange(key, 0, -1)
	if err != nil {
		t.Log("redis get range 4 err : ", err)
		return
	}
	t.Log("get range 4 success result : ", res)

	// start 如果比 end 的坐标位置大
	res, err = GetStrRange(key, -1, 0)
	if err != nil {
		t.Log("redis get range 5 err : ", err)
		return
	}
	t.Log("get range 5 success result : ", res)

}

func TestSetStrRange(t *testing.T) {
	client.Init()
	key := "strRange"
	val := "0123456789"
	result, err := client.LocalRedis.Set(key, val, 0).Result()
	if err != nil {
		t.Log("redis set err : ", err)
		return
	}
	t.Log("set success result : ", result)

	// 正下标替换
	res, err := SetStrRange(key, "abc", 0)
	if err != nil {
		t.Log("redis get range 1 err : ", err)
		return
	}
	t.Log("get range 1 success result : ", res)

	getRes, err := GetStr(key)
	if err != nil {
		t.Log("get err 1: ", err)
		return
	}
	t.Log("get result 1", getRes)

	// 负下标替换
	//res, err = SetStrRange(key, "abc", -4)
	//if err != nil {
	//	t.Log("redis set range 2 err : ", err)
	//	return
	//}
	//t.Log("set range 2 success result : ", res)

	getRes, err = GetStr(key)
	if err != nil {
		t.Log("get err 2: ", err)

		return
	}
	t.Log("get result 2", getRes)

	// 替换值过长超出最大下标
	res, err = SetStrRange(key, "abc", 9)
	if err != nil {
		t.Log("redis set range 3 err : ", err)
		return
	}
	t.Log("set range 3 success result : ", res)
	getRes, err = GetStr(key)
	if err != nil {
		t.Log("get err 3: ", err)

		return
	}
	t.Log("get result 3", getRes)

	// 替换下标超出最大下标
	res, err = SetStrRange(key, "abc", 15)
	if err != nil {
		t.Log("redis set range 4 err : ", err)
		return
	}
	t.Log("set range 4 success result : ", res)
	getRes, err = GetStr(key)
	if err != nil {
		t.Log("get err 4: ", err)

		return
	}
	t.Logf("get result 4 %#v", getRes)

}

func TestStrAppend(t *testing.T) {
	client.Init()
	key := "strAppend"
	val := "0123456789"
	result, err := client.LocalRedis.Set(key, val, 0).Result()
	if err != nil {
		t.Log("redis set err : ", err)
		return
	}
	t.Log("set success result : ", result)

	// 对某key直接追加
	res, err := StrAppend(key, "abc")
	if err != nil {
		t.Log("redis append 1 err : ", err)
		return
	}
	t.Log("append 1 success result : ", res)

	getRes, err := GetStr(key)
	if err != nil {
		t.Log("get err 1: ", err)
		return
	}
	t.Log("get result 1", getRes)

	// 处理不存在的键
	res, err = StrAppend("new_key", "abc")
	if err != nil {
		t.Log("redis append 2 err : ", err)
		return
	}
	t.Log("append 2 success result : ", res)

	getRes, err = GetStr("new_key")
	if err != nil {
		t.Log("get err 2: ", err)
		return
	}
	t.Log("get result 2", getRes)
}

func TestStrIncrByAndDecryBy(t *testing.T) {
	client.Init()
	key := "strNum"
	val := 100
	result, err := client.LocalRedis.Set(key, val, 0).Result()
	if err != nil {
		t.Log("redis set err : ", err)
		return
	}
	t.Log("set success result : ", result)

	res, err := StrIncrBy(key, 100)
	if err != nil {
		t.Log("redis incrby  err : ", err)
		return
	}
	t.Log("incrby success result : ", res)

	getRes, err := GetStr(key)
	if err != nil {
		t.Log("get err 1: ", err)
		return
	}
	t.Log("get result 1", getRes)

	res, err = StrDecrBy(key, 100)
	if err != nil {
		t.Log("redis decrby  err : ", err)
		return
	}
	t.Log("decrby success result : ", res)

	getRes, err = GetStr(key)
	if err != nil {
		t.Log("get err 2: ", err)
		return
	}
	t.Log("get result 2", getRes)

}
