package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	a := assert.New(t)

	result := Foo()
	a.Equal(0, result, "defer example Foo error, return result error")
}

func TestFoo1(t *testing.T) {
	a := assert.New(t)

	result := Foo1()
	a.Equal(2, result, "defer example Foo2 error, return result error")
}
