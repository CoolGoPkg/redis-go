package sort

import (
	"CoolGoPkg/redis-go/client"
	"fmt"
	"gopkg.in/redis.v5"
)

var StockList = []string{
	//"839680.BJ",
	//"839681.BJ",
	//"839682.BJ",
	//"839683.BJ",
	"839684.BJ",
}
var BasicInfoMap = []map[string]string{
	{"name": "Mark",
		"age":   "32",
		"score": "98.88",
	},
	{"name": "Tom",
		"age":   "18",
		"score": "93",
	},
	{"name": "Track",
		"age":   "31",
		"score": "98.9",
	},
	{"name": "Marry",
		"age":   "21",
		"score": "68.8",
	},
	{"name": "Ugly",
		"age":   "19",
		"score": "58.8",
	},
}

func PushBasicArray() {
	var err error
	for _, prodCode := range StockList {
		err = client.LocalRedis.LPush("prod_code_list", prodCode).Err()
		if err != nil {
			fmt.Println("err :", err)
			return
		}
	}
}

func HashMapSet() {
	for i, info := range StockList {
		fieldName := fmt.Sprintf("stock_%s", info)
		client.LocalRedis.HMSet(fieldName, BasicInfoMap[i])
	}
}

func GetStockRank(field string) {
	key := fmt.Sprintf("rank_stock_%s", field)
	res, err := client.LocalRedis.LRange(key, 0, -1).Result()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(res)

}

func HashMapSortByField(field string) {
	res, err := client.LocalRedis.Sort("prod_code_list", redis.Sort{
		By:    fmt.Sprintf("stock_*->%s", field),
		Get:   []string{"#", "stock_*->age", "stock_*->score"},
		Order: "DESC",
	}).Result()
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	fmt.Println(res)
}

func StoreHashMapSortByField(field string) {
	res, _ := client.LocalRedis.Sort("prod_code_list", redis.Sort{
		By:    fmt.Sprintf("stock_*->%s", field),
		Get:   []string{"#", "stock_*->age", "stock_*->score"},
		Order: "DESC",
		Store: fmt.Sprintf("rank_stock_%s", field),
	}).Result()
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return
	//}

	fmt.Println(res)
}
