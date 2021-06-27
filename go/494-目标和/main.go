package main

import "fmt"

/*
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。


示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

示例 2：
输入：nums = [1], target = 1
输出：1
*/

func dp(nums []int, target int) (int, bool) {
	if len(nums) == 1 {
		if nums[0] == target {
			return 1, true
		}

		return 0, false
	}
	var req int
	num1, ok1 := dp(nums[1:], target-nums[0])
	if ok1 {
		req += num1
	}

	num2, ok2 := dp(nums[1:], target+nums[0])
	if ok2 {
		req += num2
	}

	if !ok1 && !ok2 {
		return 0, false
	}

	return req, true
}

func findTargetSumWays(nums []int, target int) int {
	req, ok := dp(nums, target)
	if !ok {
		return 0
	}

	return req
}

func findTargetSumWays2(nums []int, target int) int {
	var req int
	var backtrack func(int, int)
	backtrack = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				req++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return req
}

func findTargetSumWays3(nums []int, target int) int {
	return 1
}

func main() {
	fmt.Print(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Print(findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3))
}
