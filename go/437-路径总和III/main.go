package main

/**
给定一个二叉树，它的每个结点都存放着一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:
1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	// 前缀结果
	prefixSum := make(map[int]int, 100)
	prefixSum[0] = 1
	return solver(root, prefixSum, sum, 0)
}

func solver(root *TreeNode, prefixSum map[int]int, target, currNum int) int {
	if root == nil {
		return 0
	}
	// 当前的结果
	res := 0
	currNum += root.Val
	// 是否有前缀结果
	res += prefixSum[currNum-target]
	// 保存但前的前缀
	prefixSum[currNum] += 1
	res += solver(root.Left, prefixSum, target, currNum)
	res += solver(root.Right, prefixSum, target, currNum)

	// 删除当前前缀结果，回退到上一层
	prefixSum[currNum] -= 1
	return res
}

func main() {

}
