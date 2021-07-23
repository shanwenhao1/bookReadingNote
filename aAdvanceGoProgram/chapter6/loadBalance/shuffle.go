package loadBalance

import (
	"fmt"
	"math/rand"
	"time"
)

var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.32:3232",
	"100.69.62.42:3232",
	"100.69.62.81:3232",
	"100.69.62.11:3232",
	"100.69.62.113:3232",
	"100.69.62.101:3232",
}

/*
	洗牌函数， fisher-yates算法, 每次随机挑选一个值放在数组末尾, 然后在n-1个元素的数组中再随机挑选一个值,
	放在数组末尾, 以此类推
	func shuffle(indexes []int) {
		for i:=len(indexes); i>0; i-- {
			lastIdx := i - 1
			idx := rand.Int(i)
			indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
		}
	}
*/
func shuffle(n int) []int {
	// rand.Perm内置了fisher-yates算法
	rand.Seed(time.Now().UnixNano())
	b := rand.Perm(n)
	return b
}

/*
	每次来新的请求时, 请求洗牌函数重新洗牌, 然后挑选第一个节点作为服务的提供者
*/
func request(params map[string]interface{}) error {
	var indexes []int
	var err error

	indexes = shuffle(7)
	maxRetryTimes := 3

	idx := indexes[0]
	fmt.Println("-------------", idx, indexes)
	for i := 0; i < maxRetryTimes; i++ {
		// TODO 转发请求至选择的服务器
		//err = apiRequest(params, endpoints[idx])
		if err == nil {
			break
		}
		idx++
	}

	if err != nil {
		// logging
		return err
	}

	return nil
}

func Request() {
	err := request(map[string]interface{}{})
	if err != nil {
		panic(err)
	}
}
