package code

import (
	"fmt"
	"os"
	"testing"
)

/*
	TestMain 用于主动执行各种测试，可以测试前后做setup和tear-down操作
*/
func TestMain(m *testing.M) {
	fmt.Println("TestMain setup.")

	retCode := m.Run() //	执行测试，包括单元测试、性能测试和示例测试

	fmt.Println("TestMain tear-down")

	os.Exit(retCode)
}
