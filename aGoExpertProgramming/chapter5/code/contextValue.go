package code

import (
	"context"
	"fmt"
	"time"
)

func HandleRequestValue(ctx context.Context, key string) {
	for {
		select {
		// 因为valueCtx不支持cancel的, 因此<-ctx.Done永远没有返回
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value(key))
			time.Sleep(2 * time.Second)
		}
	}
}

func ValueExample() {
	var key string = "parameter"
	ctxEx := InitContextExample()
	ctx := ctxEx.ExWithValue(key, "1")

	go HandleRequestValue(ctx, key)

	time.Sleep(10 * time.Second)
}
