package main

import "fmt"

/**
有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
求所能获得硬币的最大数量。

说明:
你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
示例:
输入: [3,1,5,8]
输出: 167
解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
*/
func maxCoins(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	val := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		val[i+1] = nums[i]
	}
	val[0], val[len(nums)+1] = 1, 1
	rec := make([][]int, len(nums)+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, len(nums)+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}

	return solver(0, len(nums)+1, val, rec)
}

func solver(left, right int, nums []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rec[left][right] != -1 {
		return rec[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := nums[left] * nums[i] * nums[right]
		sum += solver(left, i, nums, rec) + solver(i, right, nums, rec)
		rec[left][right] = max(sum, rec[left][right])
	}

	return rec[left][right]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	data := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(data))
}
