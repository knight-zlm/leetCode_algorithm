package main

import "fmt"

/**
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。

示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
func canJump(nums []int) bool {
	reach := 0
	n := len(nums)
	for i := 0; i <= n; i++ {
		//我们当前的位置已经超过可到达的最大位置
		if i > reach {
			return false
		}

		curReach := i + nums[i]
		if curReach > reach {
			reach = curReach
		}
	}

	return true
}

func canJump2(nums []int) bool {
	reach := 0
	n := len(nums)
	//我们最远可到达的位置是reach，并且当reach 大于等于 最后一个位置时表示判断结束
	for i := 0; i <= reach && reach < n-1; i++ {
		curReach := i + nums[i]
		if curReach > reach {
			reach = curReach
		}
	}

	return reach >= n-1
}

func main() {
	data := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(data))
}
