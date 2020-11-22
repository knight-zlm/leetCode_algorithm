package main

import "fmt"

/**
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
进阶：
你可以运用递归和迭代两种方法解决这个问题吗？
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricSlice(nodes []*TreeNode) bool {
	i := 0
	j := len(nodes) - 1
	for i < j {
		if nodes[i] == nil && nodes[j] == nil {
			i++
			j--
			continue
		}
		if nodes[i] != nil && nodes[j] != nil && nodes[i].Val == nodes[j].Val {
			i++
			j--
			continue
		}
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	// 把当前层数据放进去
	if root == nil {
		return true
	}
	nodeStack := make([]*TreeNode, 0)
	nodeStack = append(nodeStack, root)
	for len(nodeStack) != 0 {
		nextStack := make([]*TreeNode, 0)
		// 判断 是不是镜像数组
		if !isSymmetricSlice(nodeStack) {
			return false
		}
		for i := 0; i < len(nodeStack); i++ {
			// 弹出所
			curNode := nodeStack[i]
			if curNode != nil {
				nextStack = append(nextStack, curNode.Left, curNode.Right)
			}
		}
		nodeStack = nextStack
	}

	return true
}

func main() {
	data := []*TreeNode{{Val: 1}}
	fmt.Println(isSymmetricSlice(data))
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
	fmt.Println(isSymmetric(data2))
}
