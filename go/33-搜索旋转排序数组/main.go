package main

import "fmt"

/*
给你一个升序排列的整数数组 nums ，和一个整数 target 。
假设按照升序排序的数组在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：
输入：nums = [1], target = 0
输出：-1
提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4
*/

func search(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}

	return -1
}

func dp(nums []int, target, left, right int) int {
	if right < left {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	// 判断 左边 还是右边 有序，如果有序则判断target是否落在区间内
	if nums[left] < nums[mid] {
		if target > nums[left] && target < nums[mid] {
			return dp(nums, target, left+1, mid-1)
		} else {
			return dp(nums, target, mid+1, right-1)
		}
	} else {
		if target > nums[mid] && target < nums[right] {
			return dp(nums, target, mid+1, right-1)
		} else {
			return dp(nums, target, left+1, mid-1)
		}
	}
}

func search2(nums []int, target int) int {
	return dp(nums, target, 0, len(nums)-1)
}

func main() {
	data := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search2(data, 6))
}
