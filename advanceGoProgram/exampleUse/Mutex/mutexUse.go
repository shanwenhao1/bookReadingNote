package Mutex

import (
	"bookReadingNote/infra/utils"
	"fmt"
	"sync"
	"sync/atomic"
)

type mutexV struct {
	sync.Mutex
	value int
}

var totalMutex = new(mutexV)
var totalAto uint64

/*
	使用锁进行原子操作
*/
func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10000000; i++ {
		// 加锁为了保证total.value的原子性
		totalMutex.Lock()
		totalMutex.value += i
		totalMutex.Unlock()
	}
}

/*
	使用sync/atomic支持的原子操作, 效率比互斥锁高(即 worker 效率更高)
*/
func workerWithAtomic(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 10000000; i++ {
		atomic.AddUint64(&totalAto, i)
	}
}

func MutExample() {
	var wg sync.WaitGroup

	// add mutex to ensure atomic add
	wg.Add(2)
	t1 := utils.GetCurTimeUtc()
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	t2 := utils.GetCurTimeUtc()

	// use sync/atomic to ensure atomic add
	wg.Add(2)
	t3 := utils.GetCurTimeUtc()
	go workerWithAtomic(&wg)
	go workerWithAtomic(&wg)
	wg.Wait()
	t4 := utils.GetCurTimeUtc()
	fmt.Println(fmt.Sprintf("cost time: %v, get result: %v", utils.GetTimeSub(t1, t2), totalMutex.value))
	fmt.Println(fmt.Sprintf("sync/atomic cost time: %v, get result: %v", utils.GetTimeSub(t3, t4), totalAto))
}
