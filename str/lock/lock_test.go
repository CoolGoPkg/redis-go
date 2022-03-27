package lock

import (
	"CoolGoPkg/redis-go/client"
	"fmt"
	"testing"
	"time"
)

func TestLocker_Acquire(t *testing.T) {
	client.Init()
	locker := NewLocker(nil)
	if !locker.Acquire("lock") {
		locker.Release("lock")
	}
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		go DoTask(locker)
	}
}

func DoTask(locker *Locker) {
	if !locker.Acquire("lock") {
		fmt.Println("该锁被其他程序占有，无法执行任务!")
		return
	}
	ExecTask(locker)
}

func ExecTask(locker *Locker) {
	var err error
	defer func() {
		err = locker.Release("lock")
		if err != nil {
			fmt.Println("release lock err :", err)
		}
	}()
	fmt.Println("成功取锁，开始执行任务....")
	time.Sleep(1 * time.Second)
}
