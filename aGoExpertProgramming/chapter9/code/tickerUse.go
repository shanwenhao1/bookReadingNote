package code

import (
	"log"
	"time"
)

func ExampleTicker() {
	ticker := time.NewTicker(1 * time.Second) //
	defer ticker.Stop()

	// 持续从管道中获取事件
	for range ticker.C {
		log.Println("Ticker tick.")
	}
}
