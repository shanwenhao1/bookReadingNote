package singleton

import "sync"

/*
	go 单例实现示例
*/
type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	// once.Do 调用的函数只执行一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
