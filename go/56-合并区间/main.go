package main

import (
	"fmt"
	"sort"
)

/**
给出一个区间的集合，请合并所有重叠的区间。

示例 1:
输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2:
输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	mergeData := make([][]int, 0, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	mergeData = append(mergeData, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		lastIndex := len(mergeData) - 1
		lastMerge := mergeData[lastIndex]
		if cur[0] > lastMerge[1] {
			mergeData = append(mergeData, intervals[i])
			continue
		}
		if cur[1] > lastMerge[1] {
			mergeData[lastIndex][1] = cur[1]
		}
	}
	return mergeData
}

func main() {
	data := [][]int{{8, 10}, {1, 3}, {15, 20}, {2, 6}}
	fmt.Println(merge(data))
}
