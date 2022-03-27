package limit

import (
	"CoolGoPkg/redis-go/client"
	"strings"
	"testing"
	"time"
)

// 模拟一个爬虫频率限制程序
func TestLimitProgress(t *testing.T) {
	client.Init()
	crewLerLimit := NewLimiter(client.LocalRedis, "CrawlerLimit", 1, 30*time.Second)
	err := crewLerLimit.SetMaxExecuteTimes(3)
	if err != nil {
		t.Log("err : ", err)
		return
	}

	for {

		time.Sleep(5 * time.Second)
		remain, err := crewLerLimit.GetRemainExecuteTimes()
		if err != nil {
			if strings.Contains(err.Error(), "redis: nil") {
				err := crewLerLimit.SetMaxExecuteTimes(3)
				if err != nil {
					t.Log("err : ", err)
					return
				}

			} else {
				t.Log("err : ", err)
				return
			}
		}
		t.Log("Remain times : ", remain)

		if !crewLerLimit.StillValidToExecute() {
			t.Log("爬去次数受限请稍后再试。。。")
		} else {
			t.Log("爬取数据。。。")
		}
	}

}
