package main

import "fmt"

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
通过次数219,524提交次数344,471
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 把当前层数据放进去
	out := make([][]int, 0, 10)
	if root == nil {
		return out
	}
	nodeStack := make([]*TreeNode, 0)
	nodeStack = append(nodeStack, root)
	for len(nodeStack) != 0 {
		nextStack := make([]*TreeNode, 0)
		curVals := make([]int, 0, len(nodeStack)*2)
		for i := 0; i < len(nodeStack); i++ {
			// 弹出所
			curNode := nodeStack[i]
			curVals = append(curVals, curNode.Val)
			if curNode.Left != nil {
				nextStack = append(nextStack, curNode.Left)
			}
			if curNode.Right != nil {
				nextStack = append(nextStack, curNode.Right)
			}
		}
		out = append(out, curVals)
		nodeStack = nextStack
	}

	return out
}
func main() {
	data2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(levelOrder(data2))
}
