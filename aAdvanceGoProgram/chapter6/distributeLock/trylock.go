package distributeLock

import (
	"fmt"
	"sync"
)

// Lock struct of try lock
type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

// try lock, return lock result
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

// unlock
func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int

func TryLock() {
	var l = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				// log error
				println("lock failed")
				return
			}
			counter++
			fmt.Println("current counter", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}
