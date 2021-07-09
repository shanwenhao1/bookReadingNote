package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	a := assert.New(t)

	var testTree *TreeNode
	// 构造二叉树
	testTree = buildTree([][]string{{"1"}, {"2", "3"}, {"4", "#", "#", "5"}})

	result := LevelOrder(testTree)
	a.Equal([][]int{{1}, {2, 3}, {4, 5}}, result, "Level ergodic binary tree failed")
}
