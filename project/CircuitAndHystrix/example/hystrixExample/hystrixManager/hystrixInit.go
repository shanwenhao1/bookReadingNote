package hystrixManager

import (
	"fmt"
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

func HystrixRun(hy HystrixI, fakeRun RunFunc) error {
	err := hy.Run(fakeRun)
	if err != nil {
		fmt.Println("-----------", err)
	}
	return err
}

func HystrixRunWithFallback(hy HystrixI, fakeRun RunFunc, fakeFallback FallbackFunc) error {
	err := hy.RunWithFallback(fakeRun, fakeFallback)
	cirState, _, _ := hystrix.GetCircuit((hy).(*HystrixS).Name)
	if err != nil {
		if cirState.IsOpen() {
			// circuit进入Open状态时的处理
			// TODO 此处可以编写Circuit 进入Open状态的处理
			fmt.Println("----------circuit open state-----------", err)
		} else {
			// 正常请求出错的处理
			fmt.Println("------------", err)
		}
	}
	return err
}
