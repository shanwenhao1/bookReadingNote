package code

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	// 子协程2、3
	go WriteRedis(ctx)
	go WriteDataBase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running.")
			time.Sleep(2 * time.Second)

		}
	}
}

/*
	使用cancel控制子协程及孙协程的退出
*/
func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeRedis Done.")
			return
		default:
			fmt.Println("writeRedis Running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDataBase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeDatabase Done.")
			return
		default:
			fmt.Println("writeDatabase running.")
			time.Sleep(2 * time.Second)
		}
	}
}

/*
	Golang context cancelCtx 使用示例:
		由主线程控制所有其下的子孙协程全部cancel
*/
func CancelExample() {
	ctxEx := InitContextExample()
	ctx, cancel := ctxEx.ExWithCancel()
	go HandelRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutine!")
	cancel()

	//Just	for	test	whether	sub	goroutines	exit	or	no
	time.Sleep(5 * time.Second)
}
