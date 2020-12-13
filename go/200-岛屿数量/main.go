package main

import "fmt"

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
*/
func dfs(grid [][]byte, r, c int) {
	maxR := len(grid)
	maxC := len(grid[0])
	grid[r][c] = '0'
	// 上下左右
	if r-1 >= 0 && grid[r-1][c] == '1' {
		dfs(grid, r-1, c)
	}
	if r+1 < maxR && grid[r+1][c] == '1' {
		dfs(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == '1' {
		dfs(grid, r, c-1)
	}
	if c+1 < maxC && grid[r][c+1] == '1' {
		dfs(grid, r, c+1)
	}
}
func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
				continue
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1'},
		{'1'},
	}))
}
