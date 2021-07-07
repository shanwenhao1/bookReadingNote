package code

/*
	问题: 判断链表中是否有环.
		具体描述: 有环则返回true, 否则返回false, 空间复杂度为o(1)
				输入分为2部分，第一部分为链表，第二部分代表是否有环，然后回组成head头结点传入到函数里面。
				-1代表无环，其他的数字代表有环，这些参数解释仅仅是为了方便读者自测调试
		示例:
			输入：{3,2,0,-4},1
			返回值：true
			说明：第一部分{3,2,0,-4}代表一个链表，第二部分的1表示，-4到位置1，即-4->2存在一个链接，组成传入的head为一个带环的链表 ,
				 返回true
*/

/*
	解题思路: 使用快慢指针(快指针每步走二, 慢指针每步走一, 依次向后进行遍历, 如果链表中存在闭环的话, 那么快慢指针最终会相遇)
*/
func HasCycle(head *ListNode) bool {
	// write code here
	if head == nil || head.Next == nil {
		return false
	}
	p1 := head
	p2 := head.Next
	// p2 为快指针, 如果p2指向了nil 说明不存在闭环, 可退出循环
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
