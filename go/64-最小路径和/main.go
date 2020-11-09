package main

import "fmt"

/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func min(left, upper int) int {
	if left < upper {
		return left
	}
	return upper
}

func minPathSum(grid [][]int) int {
	per := make([]int, 1, len(grid[0]))
	for i := 1; i < len(grid[0]); i++ {
		per = append(per, int(^uint(0)>>1))
	}
	for i := 0; i < len(grid); i++ {
		grid[i][0] = per[0] + grid[i][0]
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = min(per[j], grid[i][j-1]) + grid[i][j]
		}
		copy(per, grid[i])
	}

	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	data := [][]int{{1, 3, 1}}
	fmt.Println(minPathSum(data))
}
