package lock

import (
	"testing"
)

func TestMutex_TryLock(t *testing.T) {
	var lock = new(Mutex)
	lock.Lock()
	defer lock.Unlock()

	if lock.TryLock() {
		t.Error("get locked Mutex, function TryLock error")
	} else {
		lock.Unlock()
		if !lock.TryLock() {
			t.Error("get unlock Mutex failed, function TryLock error")
		} else {
			t.Log("TryLock get Mutex lock")
		}
	}
}
