//go:generate mockgen -destination ../mock_hystrixManager/hystrix_mock.go bookReadingNote/project/CircuitAndHystrix/example/hystrixExample/hystrixManager HystrixI
package hystrixManager

import (
	"github.com/afex/hystrix-go/hystrix"
	"sync"
)

// 定义待执行的func, 参考hystrix包
type RunFunc func() error

// 定义fallback func
type FallbackFunc func(error) error

type HystrixI interface {
	Run(run RunFunc) error
	RunWithFallback(run RunFunc, fallback FallbackFunc) error
}

type HystrixS struct {
	Name    string                // 执行任务的名称
	config  hystrix.CommandConfig // hystrix CommandConfig
	loadMap *sync.Map
}

/*
	run: 执行的函数(熔断器监控此函数执行状态)
*/
func (s HystrixS) Run(run RunFunc) error {
	if _, ok := s.loadMap.Load(s.Name); !ok {
		s.loadMap.Store(s.Name, s.Name)
	}
	err := hystrix.Do(s.Name, func() error {
		return run()
	}, nil)
	return err
}

func (s HystrixS) RunWithFallback(run RunFunc, fallbackFunc FallbackFunc) error {
	if _, ok := s.loadMap.Load(s.Name); !ok {
		s.loadMap.Store(s.Name, s.Name)
	}
	err := hystrix.Do(s.Name, func() error {
		return run()
	}, func(err error) error {
		return fallbackFunc(err)
	})
	return err
}
