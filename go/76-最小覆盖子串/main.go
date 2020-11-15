package main

import (
	"fmt"
	"math"
)

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"

提示：
1 <= s.length, t.length <= 105
s 和 t 由英文字母组成
*/

func minWindow(s string, t string) string {
	var (
		ori = make(map[byte]int, len(t))
		cur = make(map[byte]int, len(t))
	)
	// 初始化
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}
	sLen := len(s)
	minLen := math.MaxInt32
	endL, endR := -1, -1
	distance := 0
	r, l := 0, 0
	for r < sLen {
		// 右指针项右移动找到含有t的子串
		if ori[s[r]] == 0 {
			r++
			continue
		}
		// 如果当前的字符的值的数量小于目标该字符的值则distance++
		if cur[s[r]] < ori[s[r]] {
			distance++
		}
		cur[s[r]]++
		r++
		// 检查是否包含子串
		for distance == len(t) {
			if r-l+1 < minLen {
				minLen = r - l + 1
				endL = l
				endR = r + 1
			}
			if ori[s[l]] == 0 {
				l++
				continue
			}
			if cur[s[l]] == ori[s[l]] {
				distance--
			}
			cur[s[l]]--
			l++
		}
	}
	if endL == -1 {
		return ""
	}
	fmt.Println(endR)
	return s[endL : endR-1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
