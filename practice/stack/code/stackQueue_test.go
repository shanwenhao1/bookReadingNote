package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	a := assert.New(t)

	Push(1)
	defer Pop()
	a.Equal(1, len(stack1), "Push test failed")
	a.Equal(1, stack1[0], "Push test failed, element error")
}

func TestPop(t *testing.T) {
	a := assert.New(t)

	Push(1)
	Push(2)
	a.Equal(2, len(stack1), "Push test failed")
	a.Equal(1, stack1[0], "Push test failed, element error")
	a.Equal(2, stack1[1], "Push test failed, element error")

	result := Pop()
	a.Equal(1, result, "Pop test failed, element error")
	a.Equal(0, len(stack1), "Pop test failed, push stack cleaned failed")
	a.Equal(1, len(stack2), "Pop test failed, pop stack pop element failed")
}
