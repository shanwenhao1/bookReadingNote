package code

import (
	"bookReadingNote/aGoExpertProgramming/chapter7/code"
	"fmt"
	"testing"
	"time"
)

// sub1 为子测试
func sub1(t *testing.T) {
	var a, expected = 1, 2

	actual := code.Add1(a)
	if actual != expected {
		t.Errorf("Add1(%d) = %d; expected: %d)", a, actual, expected)
	}
}

// sub2 为子测试
func sub2(t *testing.T) {
	var a, expected = 1, 3

	actual := code.Add2(a)
	if actual != expected {
		t.Errorf("Add2(%d) = %d; expected: %d)", a, actual, expected)
	}
}

// sub3 为子测试
func sub3(t *testing.T) {
	var a, expected = 1, 4

	actual := code.Add3(a)
	if actual != expected {
		t.Errorf("Add3(%d) = %d; expected: %d)", a, actual, expected)
	}
}

/*
	TestSub 内部调用sub1、sub2、sub3三个子测试
*/
func TestSub(t *testing.T) {
	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("B=3", sub3)
}

/*
	---------------------------------------------------------------------------
	子测试并发
*/
// parallelSub1 为子测试
func parallelSub1(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second) // 只是为了观察方便
	var a, expected = 1, 2

	actual := code.Add1(a)
	if actual != expected {
		t.Errorf("Add1(%d) = %d; expected: %d)", a, actual, expected)
	}
	fmt.Println("parallelSub1 out")
}

// parallelSub2 为子测试
func parallelSub2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second) // 只是为了观察方便
	var a, expected = 1, 3

	actual := code.Add2(a)
	if actual != expected {
		t.Errorf("Add2(%d) = %d; expected: %d)", a, actual, expected)
	}
	fmt.Println("parallelSub2 out")
}

// parallelSub3 为子测试
func parallelSub3(t *testing.T) {
	t.Parallel()                // 启动并发测试
	time.Sleep(1 * time.Second) // 只是为了观察方便
	var a, expected = 1, 4

	actual := code.Add3(a)
	if actual != expected {
		t.Errorf("Add3(%d) = %d; expected: %d)", a, actual, expected)
	}
	fmt.Println("parallelSub3 out")
}

/*
	TestSubParallel	通过把多个子测试放到一个组中并发执行，同时多个子测试可以共享setup和tear-down
*/
func TestSubParallel(t *testing.T) {
	// setup
	t.Logf("Setup")

	t.Run("group", func(t *testing.T) {
		t.Run("prA=1", parallelSub1)
		t.Run("prA=2", parallelSub2)
		t.Run("prB=3", parallelSub3)
	})

	// tear down
	t.Logf("teardown")
}
