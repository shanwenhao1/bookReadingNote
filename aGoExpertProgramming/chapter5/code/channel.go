package code

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	// Do some work...
	time.Sleep(time.Second)
	ch <- 1 //管道中写入一个元素表示当前协程已结束
}

/*
	创建N个channel来管理N个协程，每个协程都有一个channel用于跟父协程通信，父协程创建完所有协程中等待所有协程结束
*/
func ChannelExample() {
	channels := make([]chan int, 10) //创建一个10个元素的切片，元素类型为channel

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) //切片中放入一个channel
		go Process(channels[i])      //启动协程，传一个管道用于通信
	}

	for i, ch := range channels { //遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine", i, "quit!")
	}
}

/*
	超时测试代码模拟协程任务
*/
func timeOutGoroutine(ch chan int, sleepTime time.Duration) {
	// Do something
	// 睡眠sleepTime, 模拟任务执行时间
	time.Sleep(sleepTime)
	ch <- 1
}

/*
	主线程调用协程执行任务, 协程未在规定时间内执行完成直接退出
*/
func ChannelTimeOut(sleepTime time.Duration, timeout time.Duration) int {
	var ch chan int

	ch = make(chan int, 1)
	go timeOutGoroutine(ch, sleepTime)
	select {
	case <-ch:
		//case result := <- ch:
		//	fmt.Println(result)
		fmt.Println("call successfully!")
		return 0
	// 超时退出
	case <-time.After(time.Duration(timeout)):
		fmt.Println("timeout!")
		return -1
	}

}
