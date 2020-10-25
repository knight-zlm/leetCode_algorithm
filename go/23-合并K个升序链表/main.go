package main

import (
	"fmt"
)

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：
输入：lists = []
输出：[]
示例 3：
输入：lists = [[]]
输出：[]
提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
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

func merge2List(l1, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	curNode := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curNode.Next = l2
			l2 = l2.Next
		} else {
			curNode.Next = l1
			l1 = l1.Next
		}
		curNode = curNode.Next
	}
	if l1 != nil {
		curNode.Next = l1
	} else {
		curNode.Next = l2
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 思路 取出链表的值,两两合并的方式
	amount := len(lists)
	if amount == 0 {
		return nil
	}
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = merge2List(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}

func main() {
	var (
		l1 *ListNode
		l2 *ListNode
		l3 *ListNode
	)
	for _, val := range []int{1, 2, 4} {
		l1 = l1.Append(val)
	}
	for _, val := range []int{1, 3, 4} {
		l2 = l2.Append(val)
	}
	for _, val := range []int{2, 5} {
		l3 = l3.Append(val)
	}
	lists := []*ListNode{l1, l2, l3}
	total := mergeKLists(lists)
	total.Print()
}
