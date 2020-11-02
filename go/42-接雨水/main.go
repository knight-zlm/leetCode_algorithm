package main

import "fmt"

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
0 <= n <= 3 * 104
0 <= height[i] <= 105
*/

func dp(height []int, startIndex int) int {
	start := height[startIndex]
	totalTrap := 0
	midTrap := 0
	for i := startIndex; i < len(height); i++ {
		v := height[i]
		if v >= start {
			// 计算雨说数量,减去中间凸起消耗的空间
			tempTrap := (i-startIndex-1)*start - midTrap
			if tempTrap > 0 {
				totalTrap += tempTrap
			}
			start = v
			startIndex = i
			midTrap = 0
			continue
		}

		midTrap += v
	}

	end := startIndex
	start = height[len(height)-1]
	midTrap = 0
	for i := len(height) - 1; i >= end; i-- {
		v := height[i]
		if v >= start {
			// 计算雨说数量,减去中间凸起消耗的空间
			tempTrap := (startIndex-i-1)*start - midTrap
			if tempTrap > 0 {
				totalTrap += tempTrap
			}
			start = v
			startIndex = i
			midTrap = 0
			continue
		}

		midTrap += v
	}

	return totalTrap
}

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	return dp(height, 0)
}

func main() {
	data := []int{4, 2, 3}
	fmt.Println(trap(data))
	data = []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	fmt.Println(trap(data))
}
