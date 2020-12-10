package main

import "fmt"

/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1:
输入: [3,2,3]
输出: 3

示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
通过次数241,522提交次数370,915
*/
func majorityElement(nums []int) int {
	numMap := make(map[int]int, len(nums))
	ans := nums[0]
	ansNum := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		numMap[v]++
		if numMap[v] > ansNum {
			ans = v
			ansNum = numMap[v]
		}
	}

	return ans
}

func main() {
	fmt.Println(majorityElement([]int{6, 5, 5}))
}
