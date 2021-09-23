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
