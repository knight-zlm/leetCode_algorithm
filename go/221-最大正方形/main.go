package main

import "fmt"

/**
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4

输入：matrix = [["0","1"],["1","0"]]
输出：1

示例 3：
输入：matrix = [["0"]]
输出：0

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/
func maximalSquare(matrix [][]byte) int {
	maxSide := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide * maxSide
	}
	rows, columns := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			maxSide = max(maxSide, 1)
			curMaxSide := min(rows-i, columns-j)
			for k := 1; k < curMaxSide; k++ {
				flag := true
				if matrix[k+i][k+j] == '0' {
					break
				}
				for m := 0; m < k; m++ {
					if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
						flag = false
						break
					}
				}
				if !flag {
					break
				}
				maxSide = max(maxSide, k+1)
			}
		}
	}

	return maxSide * maxSide
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func maximalSquareEx(matrix [][]byte) int {
	maxSide := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide * maxSide
	}
	dp := make([][]int, len(matrix))
	rows, columns := len(matrix), len(matrix[0])
	// 对dp初始化
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
		for j := 0; j < columns; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if dp[i][j] == 0 {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
			maxSide = max(maxSide, dp[i][j])
		}
	}

	return maxSide * maxSide
}

func main() {
	data := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquareEx(data))
}
