package main

import (
	"bookReadingNote/project/CircuitAndHystrix/example/hystrixExample/hystrixManager"
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
)

// fake run function
func fakeRun() error {
	n := rand.Intn(5)
	if n < 3 {
		return nil
	}
	return errors.New("mock failed")
}

// fake fallback function
func fakeFallback(err error) error {
	// 此处进行时的处理
	fmt.Println("Circuit Into Open state ", err)
	return err
}

func hystrixRun(hy *hystrixManager.HystrixS) {
	err := hy.Run(fakeRun)
	if err != nil {
		fmt.Println("-----------", err)
	}
}

func hystrixRunWithFallback(hy *hystrixManager.HystrixS) {
	err := hy.RunWithFallback(fakeRun, fakeFallback)
	cirState, _, _ := hystrix.GetCircuit(hy.Name)
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
}

func main() {
	workName := "test-hystrix"
	newM := hystrixManager.NewHystrixS(workName)
	circuitStat, _, _ := hystrix.GetCircuit(workName)
	for i := 0; i < 10; i++ {
		hystrixRun(newM)
		fmt.Println(workName, " =====熔断器开启状态: ", circuitStat.IsOpen(), "请求是否允许", circuitStat.AllowRequest())
	}

	fmt.Printf("=====================================================================\n\n\n\n\n")
	workName2 := "test-hystrix-with-fallback"
	newM2 := hystrixManager.NewHystrixS(workName2)
	circuitStat2, _, _ := hystrix.GetCircuit(workName2)
	for i := 0; i < 10; i++ {
		hystrixRunWithFallback(newM2)
		fmt.Println(workName2, " =====熔断器开启状态: ", circuitStat2.IsOpen(), "请求是否允许", circuitStat2.AllowRequest())
	}
	hystrixRunWithFallback(newM2)
}
