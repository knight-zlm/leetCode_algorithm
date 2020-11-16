package main

import "fmt"

/**
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：
这是定型的图遍历问题。用深度优先算法
board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3

*/
type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist(board [][]byte, word string) bool {
	row := len(board)
	col := len(board[0])
	// 构建二维数组保存是否被访问
	vis := make([][]bool, len(board))
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, col)
	}
	// 检查函数
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		// 检查单词是否相等
		if board[i][j] != word[k] {
			return false
		}

		// 判断结束条件
		if k == len(word)-1 {
			return true
		}
		// 表示该点已经访问过了
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时要还原访问状态
		// 上下左右
		for _, d := range directions {
			newR, newC := i+d.x, j+d.y
			if 0 <= newR && newR < row && 0 <= newC && newC < col && !vis[newR][newC] {
				if check(newR, newC, k+1) {
					return true
				}
			}
		}

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	data := [][]byte{
		{'a'},
		//{'S', 'F', 'C', 'S'},
		//{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(data, "a"))
}
