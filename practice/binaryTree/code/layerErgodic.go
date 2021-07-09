package code

/*
	问题: 二叉树的层序遍历
		描述: 给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
		例如：给定的二叉树是
							3
						9		20
							15		7
			该二叉树层序遍历的结果是
			[
			[3],
			[9,20],
			[15,7]
			]
*/

/*
	解题思路: 使用队列来实现层次遍历
*/
func LevelOrder(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		curLevel := make([]int, 0)
		curNodes := nodes[:]
		nodes = make([]*TreeNode, 0)
		for _, node := range curNodes {
			// 取出当前层的节点
			curLevel = append(curLevel, node.Val)
			// 将下一层的node节点加入队列
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, curLevel)
	}
	return res
}
