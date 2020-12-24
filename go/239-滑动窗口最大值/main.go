package main

import "fmt"

/**
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

示例 3：
输入：nums = [1,-1], k = 1
输出：[1,-1]

示例 4：
输入：nums = [9,11], k = 2
输出：[11]

示例 5：
输入：nums = [4,-2], k = 2
输出：[4]
提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
*/
func maxSlidingWindow(nums []int, k int) []int {
	dangdiao := make([]int, 0, k+1)
	//dangdiao = append(dangdiao, nums[0])
	cleanDeque := func(i int) {
		for true {
			if len(dangdiao) == 0 {
				dangdiao = []int{i}
				break
			}
			if dangdiao[0] == i-k {
				dangdiao = dangdiao[1:]
			}
			if len(dangdiao) == 0 {
				dangdiao = []int{i}
				break
			}
			last := dangdiao[len(dangdiao)-1]
			if nums[last] < nums[i] {
				dangdiao = dangdiao[:len(dangdiao)-1]
			} else {
				dangdiao = append(dangdiao, i)
				break
			}
		}
	}
	for i := 0; i < k; i++ {
		cleanDeque(i)
	}
	ans := make([]int, 0, len(nums)-k+1)
	ans = append(ans, nums[dangdiao[0]])
	for i := k; i < len(nums); i++ {
		// 判断最大值的下标是不是符合预期
		cleanDeque(i)
		ans = append(ans, nums[dangdiao[0]])
	}

	return ans
}

func main() {
	data := []int{1, -1}
	fmt.Println(maxSlidingWindow(data, 1))
}
