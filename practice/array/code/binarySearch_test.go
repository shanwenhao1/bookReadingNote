package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	a := assert.New(t)

	var aN = []int{1, 2, 2, 4}
	result := Search(aN, 2)
	a.Equal(1, result, "Search failed, position error")

	aN = []int{1, 2, 2, 4, 5, 7, 9, 0}
	result = Search(aN, 5)
	a.Equal(4, result, "Search failed, position error")

	aN = []int{1, 2, 2, 4, 5, 7, 9, 0}
	result = Search(aN, 10)
	a.Equal(-1, result, "Search failed")
}
