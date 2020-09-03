package main

import (
	"bookReadingNote/project/CircuitAndHystrix/example/hystrixExample/hystrixManager"
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
)

func main() {
	// fake run function
	var fakeRun hystrixManager.RunFunc = func() error {
		n := rand.Intn(5)
		if n < 3 {
			return nil
		}
		return errors.New("mock failed")
	}
	// fake fallback function
	var fakeFallback hystrixManager.FallbackFunc = func(err error) error {
		// 此处进行时的处理
		fmt.Println("Circuit Into Open state ", err)
		return err
	}

	workName := "test-hystrix"
	newM := hystrixManager.NewHystrixS(workName)
	circuitStat, _, _ := hystrix.GetCircuit(workName)
	for i := 0; i < 10; i++ {
		hystrixManager.HystrixRun(newM, fakeRun)
		fmt.Println(workName, " =====熔断器开启状态: ", circuitStat.IsOpen(), "请求是否允许", circuitStat.AllowRequest())
	}

	fmt.Printf("=====================================================================\n\n\n\n\n")
	workName2 := "test-hystrix-with-fallback"
	newM2 := hystrixManager.NewHystrixS(workName2)
	circuitStat2, _, _ := hystrix.GetCircuit(workName2)
	for i := 0; i < 10; i++ {
		hystrixManager.HystrixRunWithFallback(newM2, fakeRun, fakeFallback)
		fmt.Println(workName2, " =====熔断器开启状态: ", circuitStat2.IsOpen(), "请求是否允许", circuitStat2.AllowRequest())
	}
}
