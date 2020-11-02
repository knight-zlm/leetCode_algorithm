package main

import "fmt"

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

func dp(nums []int, exist map[int]bool, in []int, out *[][]int) {
	for _, v := range nums {
		if _, ok := exist[v]; ok {
			continue
		}
		in = append(in, v)
		if len(in) == len(nums) {
			*out = append(*out, in)
			return
		} else {
			exist[v] = true
			newIn := make([]int, len(in), len(in)+1)
			copy(newIn, in)
			dp(nums, exist, newIn, out)
			delete(exist, v)
			in = in[:len(in)-1]
		}
	}
}

func permute(nums []int) [][]int {
	exist := make(map[int]bool, len(nums))
	in := make([]int, 0, len(nums))
	out := make([][]int, 0, len(nums)*len(nums))
	dp(nums, exist, in, &out)
	return out
}

func main() {
	data := []int{1, 2, 3}
	fmt.Println(permute(data))
}
