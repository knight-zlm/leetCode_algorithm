package main

import (
	"fmt"
	"sort"
)

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/

func dp(nums []int, cur, left, right, target int) (int, int) {
	min := int(^uint(0) >> 1)
	closestSum := target
	for left < right {
		dif := nums[cur] + nums[left] + nums[right] - target
		if dif == 0 {
			return 0, target
		} else if dif > 0 {
			if dif < min {
				min = dif
				closestSum = dif + target
			}
			right -= 1
		} else {
			if -dif < min {
				min = -dif
				closestSum = dif + target
			}
			left += 1
		}
	}
	return min, closestSum
}

func threeSumClosest(nums []int, target int) int {
	// 先排序
	// 通过双指针获取左右的值
	sort.Ints(nums)
	min := int(^uint(0) >> 1)
	closestSum := int(^uint(0) >> 1)
	for outIndex := range nums {
		//if nums[outIndex] >= target {
		//	break
		//}
		if outIndex > 0 && nums[outIndex] == nums[outIndex-1] {
			continue
		}
		dif, tempSum := dp(nums, outIndex, outIndex+1, len(nums)-1, target)
		if dif < min {
			min = dif
			closestSum = tempSum
		}
	}

	return closestSum
}

func main() {
	nums := []int{-1, 2, 1, -4}
	min := threeSumClosest(nums, 1)
	fmt.Println(min)
}
