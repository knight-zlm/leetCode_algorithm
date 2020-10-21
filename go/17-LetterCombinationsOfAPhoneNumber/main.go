package main

import "fmt"

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
2:"abc",3:"def",4:"ghi",5:"jkl",6:"mno",7:"pqrs",8:"tuv",9:"wxyz"
解题思路：
	通过递归方式拼接
*/

var disMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return disMap[digits]
	}
	words := disMap[digits[:1]]
	temp := letterCombinations(digits[1:])
	out := make([]string, 0, len(digits)*3)
	for _, i := range words {
		for _, j := range temp {
			out = append(out, i+j)
		}
	}

	return out
}

func main() {
	resp := letterCombinations("234")
	fmt.Println(resp)
}
