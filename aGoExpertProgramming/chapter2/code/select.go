package code

import "fmt"

func SelectExample() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready")
	case <-chan2:
		fmt.Println("chan2 ready")
	}

	fmt.Println("main exit.")
}
