package main

import (
	"fmt"
	"sort"
)

/**
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？
示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
提示：
0 <= nums.length <= 104
-109 <= nums[i] <= 109
*/

func longestConsecutive(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 || numsLen == 1 {
		return numsLen
	}
	// 先排序在按顺序查找
	sort.Ints(nums)
	maxLen := 0
	curLen := 1
	prev := nums[0]
	for i := 1; i < numsLen; i++ {
		if nums[i]-prev == 1 {
			curLen++
		} else if nums[i]-prev == 0 {
			continue
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = 1
		}
		prev = nums[i]
	}
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}

// 通过使用map，来通过空间换时间
func longestConsecutiveNew(nums []int) int {
	// 通过map提高计算效率
	numSet := make(map[int]bool, len(nums))
	for _, v := range nums {
		numSet[v] = true
	}
	maxLen := 0
	for v := range numSet {
		// 如果但前值的前一个值不存在，则以该值作为几点计算联系的串的长度
		if !numSet[v-1] {
			curNum := v
			curLen := 1
			// 找下一个是否存在
			for numSet[curNum+1] {
				curNum++
				curLen++
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

	return maxLen
}
func main() {
	data := []int{0, 2, 1, 1}
	fmt.Println(longestConsecutive(data))
}
