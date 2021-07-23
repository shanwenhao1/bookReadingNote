package Goroutine

import (
	"fmt"
	"sync"
)

func sayHello(s string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, " : ", i)
	}

	wg.Done()
}

func sayHello2(s string) {
	fmt.Println(s)
}

func ExGoroutine() {
	var wg = new(sync.WaitGroup)

	wg.Add(1)
	go sayHello("hello world", wg)
	for i := 0; i < 10; i++ {
		sayHello2(fmt.Sprintf("hello world-----2: %v", i))
	}

	wg.Wait()
}
