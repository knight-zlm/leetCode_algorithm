package main

import (
	"fmt"
	"sort"
)

/*
15.三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func dp(nums []int, left, right, total int) [][]int {
	out := make([][]int, 0, len(nums)/2)
	for left < right {
		dif := nums[left] + nums[right] - total
		if dif == 0 {
			out = append(out, []int{nums[left], nums[right]})
			// 去除重复的数据
			for left < right && nums[left] == nums[left+1] {
				left += 1
			}
			// 同理
			for left < right && nums[right] == nums[left-1] {
				right -= 1
			}
			left += 1
			right -= 1
		} else if dif > 0 {
			right -= 1
		} else {
			left += 1
		}
	}

	return out
}

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	out := make([][]int, 0, len(nums))
	for outIndex := range nums {
		// 如果固定位的值已经大于0，因为已经排好序了，后面的两个指针对应的值也肯定大于0，则和不可能为0，所以返回
		if nums[outIndex] > 0 {
			return out
		}
		// 排除值重复的固定位
		if outIndex > 0 && nums[outIndex] == nums[outIndex-1] {
			continue
		}

		num := nums[outIndex]
		pairs := dp(nums, outIndex+1, len(nums)-1, 0-num)
		if len(pairs) == 0 {
			continue
		}
		for _, pair := range pairs {
			temp := []int{num}
			temp = append(temp, pair...)
			out = append(out, temp)
		}
	}

	return out
}

func main() {
	//nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0, 0}
	out := threeSum(nums)
	for _, comment := range out {
		fmt.Println(comment)
	}
}
