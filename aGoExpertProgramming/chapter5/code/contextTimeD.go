package code

import "time"

/*
	Golang context timerCtx 使用示例:
		由主线程控制所有其下的子孙协程在deadline时间到期后全部cancel
*/
func TimeDeadlineExample() {
	ctxEx := InitContextExample()
	ctx, _ := ctxEx.ExWithTimeCancel(5 * time.Second)

	go HandelRequest(ctx)

	time.Sleep(10 * time.Second)
}
