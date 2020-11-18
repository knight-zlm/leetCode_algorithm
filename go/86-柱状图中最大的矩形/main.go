package main

import (
	"fmt"
)

/**
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:
输入: [2,1,5,6,2,3]
输出: 10
*/
func findMinIndex(heights []int) int {
	minV := int(^uint(0) >> 1)
	index := 0
	for i, v := range heights {
		if v < minV {
			minV = v
			index = i
		}
	}

	return index
}

func Max(data []int) int {
	maxV := ^int(^uint(0) >> 1)
	for _, v := range data {
		if v > maxV {
			maxV = v
		}
	}

	return maxV
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	// 用最小值的下标做分割
	minIndex := findMinIndex(heights)
	mRect := heights[minIndex] * len(heights)
	// 左边 和右边
	lRect := largestRectangleArea(heights[:minIndex])
	rRect := largestRectangleArea(heights[minIndex+1:])
	return Max([]int{mRect, lRect, rRect})
}

//通过单调栈解决问题

func main() {
	data := []int{2, 0, 2}
	fmt.Println(largestRectangleArea(data))
}
