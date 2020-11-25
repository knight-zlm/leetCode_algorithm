package main

import "fmt"

/**
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。

例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}
	// 前序遍历确定根节点
	rootVal := preorder[0]
	tree := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	inorderRootIndex := 0
	for i := range inorder {
		if inorder[i] == rootVal {
			inorderRootIndex = i
			break
		}
	}
	leftLen := inorderRootIndex - 0
	//rightLen := len(inorder) - inorderRootIndex - 1
	leftPreorder := preorder[1 : 1+leftLen]
	leftInorder := inorder[:inorderRootIndex]
	tree.Left = buildTree(leftPreorder, leftInorder)
	rightPreorder := preorder[1+leftLen:]
	rightInorder := inorder[inorderRootIndex+1:]
	tree.Right = buildTree(rightPreorder, rightInorder)
	return tree
}

func inorderTraversal(root *TreeNode) []int {
	var out = make([]int, 0, 500)
	if root == nil {
		return out
	}
	out = append(out, inorderTraversal(root.Left)...)
	out = append(out, root.Val)
	out = append(out, inorderTraversal(root.Right)...)
	return out
}
func main() {
	fmt.Println(inorderTraversal(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
}
