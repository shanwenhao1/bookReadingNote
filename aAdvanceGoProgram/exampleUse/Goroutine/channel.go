package Goroutine

import (
	"fmt"
	"sync"
	"time"
)

/*
	the interface example used in goroutine
*/
type ChannelI interface {
}

type ChannelS struct {
	info string
}

/*
	阻塞式的无缓冲channel
*/
func exBlockChan() {
	var (
		ch1 chan ChannelS
	)
	ch1 = make(chan ChannelS)

	go func() {
		ch1 <- ChannelS{info: "test"}
	}()

	infoC := <-ch1
	fmt.Println(infoC.info)
}

func workerChan(name string, stopCh chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-stopCh:
			fmt.Println("receive a stop signal, ", name)
			wg.Done()
			return
		default:
			fmt.Println("I am working ", name)
			time.Sleep(1 * time.Second)
		}
	}
}

/*
	用channel控制goroutine的退出
*/
func exExitChan() {
	// wg 用于控制主进程不退出
	var wg = new(sync.WaitGroup)

	stopCh := make(chan struct{})
	wg.Add(1)
	go workerChan("test name", stopCh, wg)

	time.Sleep(1)
	// 发出退出goroutine信号
	stopCh <- struct{}{}
	wg.Wait()
}

func ExChannel() {
	exBlockChan()
	exExitChan()
}
