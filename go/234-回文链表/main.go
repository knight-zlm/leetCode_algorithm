package main

import "fmt"

/**
请判断一个链表是否为回文链表。
示例 1:
输入: 1->2
输出: false
示例 2:
输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	valList := make([]int, 0, 20)
	for ; head != nil; head = head.Next {
		valList = append(valList, head.Val)
	}
	n := len(valList)
	for i := 0; i < len(valList)/2; i++ {
		if valList[i] != valList[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	data := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	fmt.Println(isPalindrome(data))
}
