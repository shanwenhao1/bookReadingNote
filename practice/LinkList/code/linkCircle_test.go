package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
	a := assert.New(t)
	var cir *ListNode
	newLink := new(ListNode)

	// 构造闭环链表
	p := newLink
	for _, v := range []int{3, 2, 0, -4} {
		q := &ListNode{
			Val: v,
		}
		p.Next = q
		p = p.Next
	}

	newLink = newLink.Next // 祛除头部的空listNode

	// 构造闭环
	m := newLink
	for {
		if m.Val == 2 {
			cir = m
		}
		if m.Next == nil {
			m.Next = cir
			break
		}
		m = m.Next
	}

	result := HasCycle(newLink)
	a.Equal(true, result, "judge linkList circle error")
}
