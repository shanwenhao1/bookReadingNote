package circuitManager

import (
	"github.com/sony/gobreaker"
	"time"
)

// self Circuit interface
type CircuitI interface {
}

type CircuitS struct {
	Manager *gobreaker.CircuitBreaker
}

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
func NewSetting(name string, reqLimit uint32) *gobreaker.Settings {
	newS := gobreaker.Settings{
		Name:        name,
		MaxRequests: reqLimit, // 处于Half-Open状态时, 允许通过的最大请求数量
		Interval:    60 * time.Second,
		Timeout:     60 * time.Second, // 进入open状态后, 经过多长时间切回Half-Open
	}
	// 设置触发进入Open状态的条件， 这里为请求超过三次失败率大于60%则触发
	newS.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 1 && failureRatio >= 0.6
	}
	return &newS
}

var CirManager *CircuitS

func init() {
	st := NewSetting("testCircuit", 50)
	newCircuits := NewCircuitS(*st)
	CirManager = newCircuits
}
