package main

import (
	"fmt"
	"sort"
)

func dp(nums []int, left, right, total int) [][]int {
	out := make([][]int, 0, len(nums)/2)
	for left < right {
		dif := nums[left] + nums[right] - total
		if dif == 0 {
			out = append(out, []int{nums[left], nums[right]})
			left += 1
		} else if dif > 0 {
			right -= 1
		} else {
			left += 1
		}
	}

	return out
}

// 三数之和
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
