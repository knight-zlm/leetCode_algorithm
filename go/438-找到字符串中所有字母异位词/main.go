package main

import "fmt"

/**
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

示例 1:
输入:
s: "cbaebabacd" p: "abc"
输出:
[0, 6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

示例 2:
输入:
s: "abab" p: "ab"
输出:
[0, 1, 2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/

func findAnagrams(s string, p string) []int {
	//
	if len(s) == 0 || len(p) == 0 {
		return []int{}
	}

	out := make([]int, 0, len(s)-1)
	subMap := make(map[byte]int, len(p))
	for _, v := range p {
		subMap[byte(v)]++
	}
	start := 0
	end := 0
	matchNum := 0
	for i := 0; i < len(s); i++ {
		if num, ok := subMap[s[end]]; ok {
			if num > 0 {
				matchNum++
			}
			subMap[s[end]]--
		}
		if end-start == len(p)-1 {
			if matchNum == len(p) {
				out = append(out, start)
			}
			if num, ok := subMap[s[start]]; ok {
				if num >= 0 {
					matchNum--
				}
				subMap[s[start]]++
			}
			start++
			end++
		} else {
			end++
		}
	}
	return out
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}
