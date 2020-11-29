package main

import (
	"fmt"
)

/**
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
示例 1：
输入：[1,2,3]

       1
      / \
     2   3
输出：6
示例 2：
输入：[-10,9,20,null,null,15,7]
   -10
   / \
  9  20
    /  \
   15   7
输出：42
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(data ...int) int {
	maxVal := ^int(^uint(0) >> 1)
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		*ans = Max(*ans, root.Val)
		return root.Val
	}
	// 左边最大值
	leftMax := Max(dfs(root.Left, ans), 0)
	// 右边最大值
	rightMax := Max(dfs(root.Right, ans), 0)
	totalMax := leftMax + rightMax + root.Val
	totalMax2 := Max(leftMax, rightMax) + root.Val
	*ans = Max(*ans, totalMax)
	// 但前节点最大贡献值
	return totalMax2
}

func maxPathSumOld(root *TreeNode) int {
	var ans int
	ans = ^int(^uint(0) >> 1)
	if root == nil {
		return 0
	}

	dfs(root, &ans)
	//if root.Val > 0 {
	//	return root.Val + ans
	//}
	return ans
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := ^int(^uint(0) >> 1)
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		//if root.Left == nil && root.Right == nil {
		//	return root.Val
		//}
		// 左边最大值
		leftMax := Max(0, dfs(root.Left))
		// 右边最大值
		rightMax := Max(0, dfs(root.Right))
		// 组合最大值
		totalMax := leftMax + rightMax + root.Val
		ans = Max(ans, totalMax)
		// 节点最大贡献值
		return Max(0, leftMax, rightMax) + root.Val
	}
	dfs(root)
	return ans
}

func main() {
	data := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -2,
			Left:  nil,
			Right: nil,
		},
		//Right: &TreeNode{
		//	Val:   3,
		//	Left:  nil,
		//	Right: nil,
		//},
	}
	fmt.Println(maxPathSumOld(data))
}
