package main

import "fmt"

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
例如，上图是一个7 x 3 的网格。有多少可能的路径？
示例 1:
输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:
输入: m = 7, n = 3
输出: 28
*/
//func dp(m int, n int, path string, uniqueMap map[string]struct{}) {
//	if n == 1 && m == 1 {
//		uniqueMap[path] = struct{}{}
//		return
//	}
//	if m == 1 {
//		for i := n; i >= 1; i-- {
//			path += "0"
//		}
//		uniqueMap[path] = struct{}{}
//		return
//	}
//	if n == 1 {
//		for i := m; i >= 1; i-- {
//			path += "1"
//		}
//		uniqueMap[path] = struct{}{}
//		return
//	}
//	if m > 1 {
//		path += "1"
//		dp(m-1, n, path, uniqueMap)
//	}
//	if n > 1 {
//		path += "0"
//		dp(m, n-1, path, uniqueMap)
//	}
//
//}

var dp = [101][101]int{}

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	} else if m == 2 && n == 2 {
		return 2
	}
	if dp[m][n] > 0 {
		return dp[m][n]
	}
	dp[m][n-1] = uniquePaths(m, n-1)
	dp[m-1][n] = uniquePaths(m-1, n)
	dp[m][n] = dp[m][n-1] + dp[m-1][n]
	return dp[m][n]
}

func main() {
	data := []int{7, 3}
	fmt.Println(uniquePaths(data[0], data[1]))
}
