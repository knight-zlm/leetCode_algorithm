package main

import "fmt"

/*
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

示例 1：
输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]

示例 2：
输入：nums = [1,1]
输出：[2]

提示：
n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
进阶：你能在不使用额外空间且时间复杂度为 O(n) 的情况下解决这个问题吗? 你可以假定返回的数组不算在额外空间内。
*/
func findDisappearedNumbers(nums []int) []int {
	total := len(nums)
	set := make(map[int]struct{}, total)
	for _, v := range nums {
		set[v] = struct{}{}
	}

	var rep []int
	for i := 1; i <= total; i++ {
		if _, ok := set[i]; !ok {
			rep = append(rep, i)
		}
	}

	return rep
}

// 原地hash
func findDisappearedNumbers2(nums []int) []int {
	for i := 0; i < len(nums); {
		if nums[i] == i+1 {
			i++
			continue
		} else {
			if nums[nums[i]-1] != nums[i] {
				nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
			} else {
				i++
			}
		}
	}

	var rep []int
	for i := range nums {
		if nums[i] != i+1 {
			rep = append(rep, i+1)
		}
	}

	return rep
}

func main() {
	fmt.Println(findDisappearedNumbers2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
