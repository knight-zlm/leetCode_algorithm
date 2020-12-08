package main

/**
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

示例 1:
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
通过次数106,454提交次数261,751
*/
func maxProduct(nums []int) int {
	maxVal, minVal, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxVal, minVal
		maxVal = max(mx*nums[i], max(mn*nums[i], nums[i]))
		minVal = min(mx*nums[i], min(mn*nums[i], nums[i]))
		ans = max(maxVal, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
