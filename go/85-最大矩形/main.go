package main

import "fmt"

/**
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积

示例 1：
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。

示例 2：
输入：matrix = []
输出：0

示例 3：
输入：matrix = [["0"]]
输出：0

示例 4：
输入：matrix = [["1"]]
输出：1

示例 5：
输入：matrix = [["0","0"]]
输出：0

提示：
rows == matrix.length
cols == matrix.length
0 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'
*/

//通过单调栈解决问题
func largestRectangleArea(heights []int) int {
	num := len(heights)
	if num == 0 {
		return 0
	}
	if num == 1 {
		return heights[0]
	}
	//添加哨兵
	newHeights := make([]int, 1, num+2)
	newHeights[0] = 0
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)
	rectArea := 0

	// 需要栈保存最后一个数据
	stack := make([]int, 1, num+2)
	stack[0] = 0
	for i := 1; i < len(newHeights); i++ {
		// 需要一直弹出知道栈内元素小于外部元素
		for newHeights[stack[len(stack)-1]] > newHeights[i] {
			//弹出栈顶元素
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 高乘长
			curArea := newHeights[topIndex] * (i - stack[len(stack)-1] - 1)
			if curArea > rectArea {
				rectArea = curArea
			}
		}
		stack = append(stack, i)
	}

	return rectArea
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([]int, len(matrix[0]))
	rectArea := 0
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] == '1' {
				dp[y] = dp[y] + 1
			} else {
				dp[y] = 0
			}
		}
		curArea := largestRectangleArea(dp)
		if rectArea < curArea {
			rectArea = curArea
		}
	}

	return rectArea
}

func main() {
	//data := [][]byte{
	//	[]byte("10100"),
	//	[]byte("10111"),
	//	[]byte("11111"),
	//	[]byte("10010"),
	//}
	data := [][]byte{}
	fmt.Println(maximalRectangle(data))
}
