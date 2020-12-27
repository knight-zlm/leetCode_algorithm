package main

import (
	"fmt"
	"math"
)

/**
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.

示例 2:
输入: n = 13
输出: 2
解释: 13 = 4 + 9.
*/
func numSquares(n int) int {
	squareNum := make([]int, 0, n/2)
	for i := 0; i <= int(math.Sqrt(float64(n))); i++ {
		squareNum = append(squareNum, i*i)
	}
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = n
	}
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		for _, j := range squareNum {
			if i < j {
				break
			}
			dp[i] = min(dp[i], dp[i-j]+1)
		}
	}
	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(numSquares(12))
}
