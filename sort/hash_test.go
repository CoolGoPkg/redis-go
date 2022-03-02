package sort

import (
	"CoolGoPkg/redis-go/client"
	"testing"
)

func init() {
	client.Init()
}

func InitBasicInfo() {
	PushBasicArray()
	HashMapSet()
}

func TestHashMapSortByField(t *testing.T) {
	HashMapSortByField("age")
}

func TestStoreHashMapSortByField(t *testing.T) {
	StoreHashMapSortByField("age")
}

func TestGetStockRank(t *testing.T) {
	GetStockRank("age")
}
