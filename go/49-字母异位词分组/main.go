package main

import (
	"fmt"
	"sort"
)

/**
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

说明：
所有输入均为小写字母。
不考虑答案输出的顺序。
*/
func groupAnagrams(strs []string) [][]string {
	mapTo := make(map[string][]string, 0)
	for i := 0; i < len(strs); i++ {
		temp := strs[i]
		newSlice := []byte(temp)
		sort.SliceStable(newSlice, func(i, j int) bool {
			return newSlice[i] < newSlice[j]
		})
		key := string(newSlice)
		mapTo[key] = append(mapTo[key], temp)
	}

	out := make([][]string, 0, len(strs))
	for _, v := range mapTo {
		out = append(out, v)
	}

	return out
}

func main() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	data3 := groupAnagrams(data)
	fmt.Println(data3)
}
