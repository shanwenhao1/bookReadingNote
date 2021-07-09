package code

import "fmt"

/*
	问题: 分别按照二叉树先序，中序和后序打印所有的节点。
		输入->{1,2,3}
		返回<-[[1,2,3],[2,1,3],[2,3,1]]
*/
var pres, ins, posts []int

func threeOrders(root *TreeNode) [][]int {
	// write code here
	pres, ins, posts = []int{}, []int{}, []int{}
	find(root)
	return [][]int{pres, ins, posts}
}
func find(node *TreeNode) {
	if node != nil {
		pres = append(pres, node.Val)
		find(node.Left)
		ins = append(ins, node.Val)
		find(node.Right)
		posts = append(posts, node.Val)
	}
}

func NormalErRun() {
	var tre1 TreeNode
	tre1.Val = 1
	tre1.Left = &TreeNode{
		2,
		&TreeNode{
			4,
			nil,
			nil,
		},
		nil,
	}
	tre1.Right = &TreeNode{
		3,
		nil,
		nil,
	}
	result := threeOrders(&tre1)
	fmt.Println(result)
}
