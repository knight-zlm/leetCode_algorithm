package main

import "fmt"

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
示例 2:
输入: [4,1,2,1,2]
输出: 4
*/
func singleNumber(nums []int) int {
	numSet := make(map[int]bool, len(nums)/2+1)
	for _, v := range nums {
		if numSet[v] {
			delete(numSet, v)
		} else {
			numSet[v] = true
		}
	}
	cur := nums[0]
	for k := range numSet {
		return k
	}
	return cur
}

func singleNumberNew(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	data := []int{2, 2, 1}
	fmt.Println(singleNumber(data))
}
