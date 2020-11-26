package main

/**
给定一个二叉树，原地将它展开为一个单链表。

例如，给定二叉树
    1
   / \
  2   5
 / \   \
3   4   6

将其展开为：
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0, 10)
	if root == nil {
		return list
	}
	list = append(list, root)
	list = append(list, dfs(root.Left)...)
	list = append(list, dfs(root.Right)...)

	return list
}

func flatten(root *TreeNode) {
	list := dfs(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}
func main() {

}
