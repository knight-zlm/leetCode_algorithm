package main

import "fmt"

/*
合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Append(val int) *ListNode {
	if l == nil {
		return &ListNode{Val: val}
	}
	head := l
	for head.Next != nil {
		head = head.Next
	}
	head.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	return l
}

func (l *ListNode) Print() {
	out := ""
	head := l
	if head == nil {
		fmt.Println("nil")
		return
	}
	if head.Next == nil {
		fmt.Println(l.Val)
		return
	}

	for head.Next != nil {
		out += fmt.Sprintf("%d->", head.Val)
		head = head.Next
	}
	out += fmt.Sprintf("%d", head.Val)
	fmt.Println(out)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{}
	nextNew := newHead
	for l1 != nil && l2 != nil {
		nextNew.Next = &ListNode{}
		nextNew = nextNew.Next
		if l1.Val > l2.Val {
			nextNew.Val = l2.Val
			l2 = l2.Next
		} else {
			nextNew.Val = l1.Val
			l1 = l1.Next
		}
	}
	if l1 != nil {
		nextNew.Next = l1
	}
	if l2 != nil {
		nextNew.Next = l2
	}

	return newHead.Next
}

func main() {
	var l1 *ListNode
	var l2 *ListNode
	for _, val := range []int{1, 2, 4} {
		l1 = l1.Append(val)
	}
	for _, val := range []int{1, 3, 4} {
		l2 = l2.Append(val)
	}
	l3 := mergeTwoLists(l1, l2)
	l3.Print()
}
