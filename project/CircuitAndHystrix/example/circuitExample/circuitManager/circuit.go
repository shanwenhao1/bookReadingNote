package circuitManager

import (
	"errors"
	"fmt"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// self Circuit interface
type CircuitI interface {
	Get(url string) ([]byte, error)
	fakeGet() (interface{}, error)
	FakeMany()
}

type CircuitS struct {
	Manager *gobreaker.CircuitBreaker
}

func (cir *CircuitS) Get(url string) ([]byte, error) {
	body, err := cir.Manager.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})

	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func (cir *CircuitS) fakeGet() (interface{}, error) {
	time.Sleep(time.Second)
	n := rand.Intn(5)
	if n < 3 {
		return "success", nil
	}
	return nil, errors.New("failed")
}

func (cir *CircuitS) FakeMany() {
	var wg sync.WaitGroup
	// 成功计数
	succeed := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		time.Sleep(time.Second)
		go func(i int, succeed *int) (interface{}, error) {
			wg.Done()
			fmt.Println("----------start mock execute goroutines: ", i)
			// 说明, 当breaker状态进入Open后, cir.Manager.Execute直接返回错误, 请求进入Open状态,
			body, err := cir.Manager.Execute(func() (interface{}, error) {
				// -----------------------------------------------------------------------------
				// open状态下不执行, 直接返回
				resp, err := cir.fakeGet()
				fmt.Println("circuit normal, state: ", cir.Manager.State())
				if err != nil {
					return nil, err
				} else {
					*succeed += 1
				}
				fmt.Println(fmt.Sprintf("----------goroutines %d result is: ", i), resp, err, "succeed number: ", *succeed)
				return resp, nil
				// ------------------------------------------------------------------------------
			})
			if err != nil {
				fmt.Println("Into Open state: ", cir.Manager.State(), "result: ", body, err)
				return nil, err
			}
			return body, nil
		}(i, &succeed)
	}
	wg.Wait()
	time.Sleep(time.Second * 20)
}
