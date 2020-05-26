package main

import (
	"advanceGoProgramming/advanceGoProgram/chapter6/worm"
	"context"
	"fmt"
)

func main() {
	//cgo.BaseCgo()
	//chapter1.RunFunc()
	//chapter1.ContextRun()
	//cgo.BaseCgo()
	//cgo2.Hello()
	//cgo3.Hello()
	//cgoFunc.FuncCgo()
	//sort.SortRun()
	//sort2.SortRun()
	//cToGo.MemoryRun()
	//goToC.MemoryRun()
	//cgoM.Run()
	//pkg.AssemblyRun()
	//pkg.Hello()
	//begin.HelloClient()
	//rpcSafe.HelloClient()
	//rpcSafe.HelloJsonClient()
	//rpcWatch.WatchClient()

	/*
		chapter 5
	*/
	//baseCgi.BaseServer()
	//baseCgi.BaseMiddleware()
	//baseCgi.TestValidate()
	//baseCgi.ValidateRun()
	//tokenLimit.BucketRun()
	//hashFunc.MurmurHash()

	/*(
	chapter 6
	*/
	//distributedId.UseSnowFlake()
	//distributedId.UseSonyFlake()
	//distributeLock.TryLock()
	//loadBalance.Request()
	//loadBalance.ShuffleCompare()
	worm.SingleWorm()
}

func contextUse() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}
}
