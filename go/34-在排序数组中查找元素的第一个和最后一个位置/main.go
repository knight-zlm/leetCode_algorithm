package main

import "fmt"

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。

示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func dp(nums []int, target, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) >> 1
	midVal := nums[mid]
	if midVal == target {
		return mid
	} else if midVal > target {
		return dp(nums, target, start, mid-1)
	} else {
		return dp(nums, target, mid+1, end)
	}
}

func leftDp(nums []int, target, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) >> 1
	midVal := nums[mid]
	if midVal == target {
		if mid == start || nums[mid-1] != target {
			return mid
		}

		return leftDp(nums, target, start, mid-1)
	}
	if midVal > target {
		return leftDp(nums, target, start, mid-1)
	} else {
		return leftDp(nums, target, mid+1, end)
	}
}

func RightDp(nums []int, target, start, end int) int {
	if end < start {
		return -1
	}
	mid := (start + end) >> 1
	midVal := nums[mid]
	if midVal == target {
		if mid == end || nums[mid+1] != target {
			return mid
		}

		return RightDp(nums, target, mid+1, end)
	}
	if midVal > target {
		return RightDp(nums, target, start, mid-1)
	} else {
		return RightDp(nums, target, mid+1, end)
	}
}

func searchRange(nums []int, target int) []int {
	findIndex := dp(nums, target, 0, len(nums)-1)
	if findIndex == -1 {
		return []int{-1, -1}
	}
	left := -1
	right := -1
	if findIndex > 0 && nums[findIndex-1] == target {
		left = leftDp(nums, target, 0, findIndex-1)
	} else {
		left = findIndex
	}
	if findIndex < len(nums)-1 && nums[findIndex+1] == target {
		right = RightDp(nums, target, findIndex+1, len(nums)-1)
	} else {
		right = findIndex
	}

	return []int{left, right}
}

func main() {
	data := []int{0, 0, 0, 2, 2, 2, 2, 2, 2, 3, 4, 5, 5, 5, 6, 7, 9, 9, 9, 9, 10, 11, 12}
	fmt.Println(searchRange(data, 9))
}
