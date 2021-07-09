package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	a := assert.New(t)

	testTree := buildTree([][]string{{"3"}, {"5", "1"}, {"6", "2", "0", "8"}, {"#", "#", "7", "4", "#", "#", "#", "#"}})
	result := LowestCommonAncestor(testTree, 5, 1)
	a.Equal(3, result, "test LowestCommonAncestor failed, result error")

	result = LowestCommonAncestor(testTree, 7, 8)
	a.Equal(3, result, "test LowestCommonAncestor failed, result error")

	result = LowestCommonAncestor(testTree, 7, 6)
	a.Equal(5, result, "test LowestCommonAncestor failed, result error")

	/*
		未考虑元素不存在边界情况
		result = LowestCommonAncestor(testTree, 3, 9)
		a.Equal(-1, result, "test LowestCommonAncestor failed, result error")
	*/
}
