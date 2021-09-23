package code

import (
	"fmt"
	"sync"
	"time"
)

func exampleGoroutine(wg *sync.WaitGroup, number uint8) {
	defer wg.Done() // goroutine执行结束后将计数器减1, 最好使用defer执行

	time.Sleep(time.Second)

	fmt.Println(fmt.Sprintf("Goroutine %d finished!", number))
}

func WaitGroupExample() {
	var wg sync.WaitGroup

	wg.Add(2) // 设置计数器, 与启动的goroutine的个数相同
	// 协程1
	go exampleGoroutine(&wg, 1)
	// 协程2
	go exampleGoroutine(&wg, 2)
	wg.Wait() //主goroutine阻塞等待计数器变为0
	fmt.Println("All Goroutine finished!")
}
