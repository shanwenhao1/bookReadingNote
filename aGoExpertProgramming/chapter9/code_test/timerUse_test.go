package code_test

import (
	"bookReadingNote/aGoExpertProgramming/chapter9/code"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExampleTimer(t *testing.T) {
	a := assert.New(t)
	exChan := make(chan string, 1)

	result := code.ExampleTimer(exChan)
	// 未通过管道传入信号, 因此会触发定时器Timer
	a.Equal(false, result, "timer example error, timeout!")

	// 先给channel缓存发送信号(这里为了方便演示, 正常应该使用channel获取协程返回信息的设计方式)
	exChan <- "excellent normal example"
	result = code.ExampleTimer(exChan)
	a.Equal(true, result, "timer not trigger")
}
