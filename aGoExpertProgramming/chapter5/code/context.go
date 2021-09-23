package code

import (
	"context"
	"time"
)

type ContextExample struct {
	ContextBackground context.Context
}

func (ctxE ContextExample) ExWithCancel() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctxE.ContextBackground)
	return ctx, cancel
}

func (ctxE ContextExample) ExWithTimeCancel(duration time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctxE.ContextBackground, duration)
	return ctx, cancel
}

func (ctxE ContextExample) ExWithValue(key interface{}, value interface{}) context.Context {
	ctx := context.WithValue(ctxE.ContextBackground, key, value)
	return ctx
}

func InitContextExample() *ContextExample {
	var background = context.Background() // 	context.Background()返回context自定义的空context
	/* context提供的四种不同类型的context的创建方法
	WithCancel()
	WithDeadline()c
	WithTimeout()
	WithValue()
	*/
	var ctxEx = new(ContextExample)
	ctxEx.ContextBackground = background

	return ctxEx
}
