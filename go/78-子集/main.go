package main

import (
	"fmt"
)

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
示例:
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
func subsets(nums []int) [][]int {
	out := make([][]int, 0, len(nums))
	out = append(out, []int{})
	for _, v := range nums {
		new := make([][]int, len(out))
		// 深拷贝的问题
		for copyIndex := range out {
			new[copyIndex] = make([]int, len(out[copyIndex]), len(out[copyIndex])+1)
			copy(new[copyIndex], out[copyIndex])
			new[copyIndex] = append(new[copyIndex], v)
		}

		out = append(out, new...)
	}

	return out
}

func main() {
	data := []int{9, 0, 3, 4, 7}
	fmt.Println(subsets(data))
}
