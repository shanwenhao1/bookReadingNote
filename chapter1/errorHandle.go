package chapter1

import (
	"fmt"
	"github.com/gogo/protobuf/types"
)

func ParseJson(input string) (s *types.Syntax, err error) {
	// 捕获JSON解析器异常并打印错误信息
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("JSON: internal error: %v", p)
		}
	}()
	// 开始解析工作
	return s, err
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}

/*
	自定义Error接口类型, 为error接口类型的扩展
		给错误增加调用栈信息		支持错误的多级嵌套包装			支持错误码格式
*/
type Error interface {
	Caller() []CallerInfo
	Wraped() []error
	Code() int
	error

	Private()
}
