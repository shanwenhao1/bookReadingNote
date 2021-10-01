package code

import (
	"fmt"
	"time"
)

func ExampleTimer(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second) // 定时任务, 1秒钟后发送信号

	select {
	case <-conn: // 收到主动信号, 停止timer定时任务
		timer.Stop()
		return true
	case <-timer.C: // 一秒中后超时, timer会自动向timer管道中写入数据
		fmt.Println("WaitChannel timeout!")
		return false
	}
}
