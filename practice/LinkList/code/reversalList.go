package code

import "fmt"

/*
	问题: 输入一个链表, 反转链表后, 输出新链表的表头
*/

func reverseList(pHead *ListNode) *ListNode {
	var (
		newHead *ListNode
		midHead *ListNode
	)

	for {
		if pHead == nil {
			break
		}
		midHead = pHead
		pHead = pHead.Next
		midHead.Next = newHead
		newHead = midHead
	}
	return newHead
}

func ReversalRun() {
	var newL = new(ListNode)
	p := newL
	for _, v := range []int{1, 3, 2, 4, 5} {
		q := &ListNode{Val: v}
		p.Next = q
		p = p.Next
	}
	newL = newL.Next

	result := reverseList(newL)
	fmt.Println("--------Reverse Linklist result: ")
	for p := result; p != nil; p = p.Next {
		fmt.Print(p.Val, ", ")
	}
}
