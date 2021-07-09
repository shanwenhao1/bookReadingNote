package code

/*
	问题: 二叉树中找到两个节点的最近公共祖先节点
		描述: 给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
			 (注：本题保证二叉树中每个节点的val值均不相同)
			示例:
				输入：[3,5,1,6,2,0,8,#,#,7,4],5,1
				返回值：3


*/

/*
	解题思路：dfs 深度遍历. 有三种情况
		- 1. o1, o2分别在公共祖先左右两侧
		- 2. 祖先是o1, o2 在祖先左/右侧
		- 3. 祖先是o2, o1 在祖先左/右侧
*/
func LowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	parent := findPar(root, o1, o2)
	if parent == nil {
		return -1
	}
	return parent.Val
}

func findPar(root *TreeNode, o1 int, o2 int) *TreeNode {
	// 如果找到o1或者o2、或者节点为空时则直接返回
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	// 当前节点左子树中查找
	left := findPar(root.Left, o1, o2)
	// 当前节点右子树中查找
	right := findPar(root.Right, o1, o2)
	// 当左子树为空时, 说明o1、o2均不在当前左子树中, 返回右子树继续递归
	if left == nil {
		return right
	} else if right == nil {
		// 同理, 当右子树为空时, 说明o1, o2均不在右子树中, 返回左子树中继续递归
		return left
	} else {
		// 当左右子树都不为空时, 说明两个节点分别坐落于left和right中, 此时的节点即为最近parent节点
		return root
	}
}
