package main

import (
	"fmt"
	"sort"
)

/**
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
进阶：
你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]

提示：
链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	//temp, temp1, temp2 := dummyHead, head1, head2
	temp := dummyHead
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			temp.Next = head1
			head1 = head1.Next
		} else {
			temp.Next = head2
			head2 = head2.Next
		}
		temp = temp.Next
	}
	if head1 != nil {
		temp.Next = head1
	}
	if head2 != nil {
		temp.Next = head2
	}

	return dummyHead.Next
}

func mergerSort(head, tail *ListNode) *ListNode {
	// 快慢指针
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	// 找到中间点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(mergerSort(head, mid), mergerSort(mid, tail))
}

func sortListNew(head *ListNode) *ListNode {
	return mergerSort(head, nil)
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head
	vals := make([]int, 0, 1000)
	for temp != nil {
		vals = append(vals, temp.Val)
		temp = temp.Next
	}
	sort.Ints(vals)
	newHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	sen := newHead
	for _, v := range vals {
		sen.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		sen = sen.Next
	}
	return newHead.Next
}

func main() {
	data := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	resp := sortListNew(data)
	fmt.Println(resp)
}
