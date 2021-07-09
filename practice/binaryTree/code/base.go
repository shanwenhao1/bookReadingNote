package code

import "strconv"

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createTree(node []*TreeNode, appendNode []string) []*TreeNode {
	var nextLayerRoot []*TreeNode
	for pos, curNode := range node {
		var leftNode, rightNode *TreeNode
		if appendNode[2*pos] != "#" {
			num, _ := strconv.Atoi(appendNode[2*pos])
			leftNode = &TreeNode{
				Val: num,
			}
		} else {
			leftNode = nil
		}

		if appendNode[2*pos+1] != "#" {
			num, _ := strconv.Atoi(appendNode[2*pos+1])
			rightNode = &TreeNode{
				Val: num,
			}
		} else {
			rightNode = nil
		}
		if curNode == nil {
			continue
		}
		curNode.Left = leftNode
		curNode.Right = rightNode
		nextLayerRoot = append(nextLayerRoot, leftNode)
		nextLayerRoot = append(nextLayerRoot, rightNode)
	}
	return nextLayerRoot
}

func buildTree(a [][]string) *TreeNode {
	var curRoot []*TreeNode
	var root = &TreeNode{}
	for _, val := range a {
		// 根节点
		if len(val) == 1 {
			num, _ := strconv.Atoi(val[0])
			root.Val = num
			curRoot = append(curRoot, root)
			continue
		}
		curRoot = createTree(curRoot, val)
	}
	return root
}
