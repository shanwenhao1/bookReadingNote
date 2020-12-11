package code

import (
	"fmt"
	"sort"
)

/*
 * 给定一个无序单链表，实现单链表的排序(按升序排序)。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 升序排序
func (list *ListNode) OrderAsc() *ListNode {
	var tmp = []int{}
	for p := list; p != nil; p = p.Next {
		tmp = append(tmp, p.Val)
	}
	sort.Ints(tmp)
	// 重新分配了内存空间, 新建了一个ListNode
	result := new(ListNode)
	p := result
	for _, val := range tmp {
		q := &ListNode{Val: val}
		p.Next = q
		p = p.Next
	}
	return result.Next
}

/**
 *
 * @param head ListNode类 the head node
 * @return ListNode类
 */
func sortInList(head *ListNode) *ListNode {
	// write code here
	return head.OrderAsc()
}

func SortRun() {
	var newL = new(ListNode)
	p := newL
	for _, v := range []int{1, 3, 2, 4, 5} {
		q := &ListNode{Val: v}
		p.Next = q
		p = p.Next
	}
	newL = newL.Next

	result := newL.OrderAsc()
	fmt.Println("--------Sort Linklist result: ")
	for p := result; p != nil; p = p.Next {
		fmt.Print(p.Val, ", ")
	}
}
