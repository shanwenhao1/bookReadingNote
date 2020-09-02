package hystrixManager

import (
	"github.com/afex/hystrix-go/hystrix"
	"sync"
)

func NewHystrixS(name string) *HystrixS {
	var config = hystrix.CommandConfig{
		Timeout:                1000,  // 执行command的超时时间(毫秒)
		MaxConcurrentRequests:  10,    // command的最大并发量
		RequestVolumeThreshold: 5,     // 触发熔断器判断前, 请求不得低于此数, 否则不给予判断
		SleepWindow:            10000, // 熔断触发后, 过多长时间, 进行熔断器是否开启检测
		ErrorPercentThreshold:  20,    // 错误率, 达到该错误率后触发熔断
	}
	// 为config 命名
	hystrix.ConfigureCommand(name, config)
	hyS := HystrixS{
		Name:    name,
		config:  config,
		loadMap: new(sync.Map),
	}
	return &hyS
}
