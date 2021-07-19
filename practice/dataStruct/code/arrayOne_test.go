package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	a := assert.New(t)

	result := SingleNumber([]int{2, 2, 1})
	a.Equal(1, result, "SingleNumber error")

	result = SingleNumber([]int{2, 2, 1, 3, 3})
	a.Equal(1, result, "SingleNumber error")
}
