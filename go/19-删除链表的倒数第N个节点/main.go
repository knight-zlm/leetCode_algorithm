package main

import "fmt"

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Append(val int) *ListNode {
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 需要两个指针a和b a先移动n步，然后a和b一起移动，a到达最后时，b就是要删除的节点
	// 所以我们需要保存 b节点的父节点
	before := head
	after := head
	for i := 0; i < n; i++ {
		before = before.Next
	}
	var father *ListNode
	for before != nil {
		before = before.Next
		father = after
		after = after.Next
	}
	if father == nil {
		return after.Next
	}
	// 删除节点
	father.Next = after.Next
	return head
}

/**
别人的优秀代码
一次遍历: 需要使用快慢针 来做定位;
	当快针到达结果的时候;
	慢针正好在 剔除 节点的前节点
注意: 使用哨兵节点防止出现删除首节点的问题
*/
func removeNthFromEndEx(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 增加虚拟头节点:哨兵节点; 如果不用; 当删除的节点为 head 节点时候,就会出现异常
	//     总结: 凡是删除节点的操作,都需要引入哨兵节点
	dummyHead := &ListNode{Next: head}

	quickNode := dummyHead
	slowNode := dummyHead

	// 快指针的Next == nil 则说 快指针已经到了 最后一个节点,此时 slowNode 为要删除的节点的前一个节点
	for quickNode.Next != nil {
		quickNode = quickNode.Next
		if n == 0 {
			slowNode = slowNode.Next
		} else {
			n--
		}
	}
	// fmt.Printf("慢指针: %v 快指针: %v\n", slowNode.Val, quickNode.Val)
	slowNode.Next = slowNode.Next.Next
	return dummyHead.Next
}

func main() {
	head := &ListNode{
		Val: 1,
	}
	//for _, value := range []int{2} {
	//	head.Append(value)
	//}
	head.Print()
	head = removeNthFromEnd(head, 1)
	head.Print()
}
