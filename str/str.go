package str

import (
	"CoolGoPkg/redis-go/client"
	"encoding/json"
	"fmt"
)

// SetStr doc
func SetStr(key, val string) error {
	res, err := client.LocalRedis.Set(key, val, 0).Result()
	if err != nil {
		return err
	}

	fmt.Println("set key result : ", res)
	return nil

}

// GetStr doc
func GetStr(key string) (string, error) {
	return client.LocalRedis.Get(key).Result()
}

// SetStrNX redis set NX=true
func SetStrNX(key, val string) error {
	sec, err := client.LocalRedis.SetNX(key, val, 0).Result()
	if err != nil {
		return err
	}

	if !sec {
		return fmt.Errorf("the key : %s has been setted", key)
	}

	return nil
}

// SetStrXX redis set XX=true
func SetStrXX(key, val string) error {
	sec, err := client.LocalRedis.SetXX(key, val, 0).Result()
	if err != nil {
		return err
	}

	if !sec {
		return fmt.Errorf("the key : %s is not find", key)
	}

	return nil
}

// GetSetStr  doc
func GetSetStr(key, val string) (string, error) {
	return client.LocalRedis.GetSet(key, val).Result()
}

type Student struct {
	Name string
	Age  int
}

func (s Student) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

// MSetStr doc
func MSetStr() (string, error) {
	s := Student{"mark", 18}
	return client.LocalRedis.MSet("key6", s).Result()
}

// MGetStr doc
func MGetStr() ([]interface{}, error) {
	return client.LocalRedis.MGet("key1", "key2", "key3", "key4", "key5", "key6").Result()
}

// MSetNXStr doc
func MSetNXStr() (bool, error) {
	return client.LocalRedis.MSetNX("a", "a", "b", "b").Result()
}

// GetStrLen doc
func GetStrLen(key string) (int64, error) {
	return client.LocalRedis.StrLen(key).Result()
}

// GetStrRange doc
func GetStrRange(key string, start, end int64) (string, error) {
	return client.LocalRedis.GetRange(key, start, end).Result()
}

// SetStrRange doc
func SetStrRange(key, val string, offset int64) (int64, error) {
	return client.LocalRedis.SetRange(key, offset, val).Result()
}

// StrAppend doc
func StrAppend(key, val string) (int64, error) {
	return client.LocalRedis.Append(key, val).Result()
}

// StrIncrBy doc
func StrIncrBy(key string, val int64) (int64, error) {
	return client.LocalRedis.IncrBy(key, val).Result()
}

// StrDecrBy doc
func StrDecrBy(key string, val int64) (int64, error) {
	return client.LocalRedis.DecrBy(key, val).Result()
}

// StrIncr doc
func StrIncr(key string) (int64, error) {
	return client.LocalRedis.Incr(key).Result()
}

// StrDecr doc
func StrDecr(key string) (int64, error) {
	return client.LocalRedis.Decr(key).Result()
}

// StrIncrByFloat doc
func StrIncrByFloat(key string, val float64) (float64, error) {
	return client.LocalRedis.IncrByFloat(key, val).Result()
}
