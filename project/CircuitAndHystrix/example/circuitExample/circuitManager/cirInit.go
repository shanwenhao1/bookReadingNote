package circuitManager

import (
	"github.com/sony/gobreaker"
	"time"
)

var CirManager *CircuitS

/*
	generate a new CircuitS instance
	setS: the gobreaker.Settings, include circuit settings
*/
func NewCircuitS(setS gobreaker.Settings) *CircuitS {
	newCirBreaker := gobreaker.NewCircuitBreaker(setS)
	newCircuitS := CircuitS{Manager: newCirBreaker}
	return &newCircuitS
}

/*
	generate the circuits settings
*/
func NewSetting(name string, reqLimit uint32) gobreaker.Settings {
	newS := gobreaker.Settings{
		Name:        name,
		MaxRequests: reqLimit,         // 处于Half-Open状态时, 允许通过的最大请求数量
		Interval:    0,                // 熔断器关闭状态下定期清除Counts计数的循环周期, 如果Interval为0，则close状态下不会清除计数, 一般为 60 * time.Second这种形式
		Timeout:     60 * time.Second, // 进入open状态后, 经过多长时间切回Half-Open
	}
	// 设置触发进入Open状态的条件， 这里为请求超过3次,失败次数大于2次且失败率大于20%则触发
	newS.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && counts.TotalFailures >= 2 && failureRatio >= 0.2
	}
	return newS
}

func init() {
	st := NewSetting("testCircuit", 5)
	newCircuits := NewCircuitS(st)
	CirManager = newCircuits
}
