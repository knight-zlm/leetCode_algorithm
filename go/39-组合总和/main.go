package main

import (
	"fmt"
	"sort"
)

/*
*给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。

示例 1：
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]

示例 2：
输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

提示：
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/

func dp(candidates []int, curElements []int, target int, out *[][]int) {
	for i, v := range candidates {
		if v == target {
			curElements = append(curElements, v)
			*out = append(*out, curElements)
			return
		} else if v < target {
			//数组拷贝问题一定要注意
			newRes := make([]int, len(curElements), len(curElements)+1)
			copy(newRes, curElements)
			newRes = append(newRes, v)
			dp(candidates[i:], newRes, target-v, out)
		} else {
			return
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	// 先排序
	sort.Ints(candidates)
	out := make([][]int, 0, len(candidates))
	for i, v := range candidates {
		if v > target {
			return out
		} else if v == target {
			out = append(out, []int{v})
			return out
		} else {
			curElements := []int{v}
			dp(candidates[i:], curElements, target-v, &out)
		}
	}

	return out
}

func main() {
	data := []int{2, 7, 6, 3, 5, 1}
	resp := combinationSum(data, 9)
	fmt.Println(resp)
	data = []int{2, 3, 6, 7}
	resp = combinationSum(data, 7)
	fmt.Println(resp)
}
