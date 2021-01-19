package main

import (
	"fmt"
	"strconv"
)

/**
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/
func decodeString(s string) string {
	// 重复次数的栈
	repeatStk := make([]int, 0, len(s))
	// 字符开始替换位置
	leftIndexes := make([]int, 0, len(s))
	tempNum := ""
	for i := 0; i < len(s); i++ {
		if 48 <= uint(s[i]) && uint(s[i]) <= 57 {
			tempNum += string(s[i])
			continue
		}
		if s[i] == '[' {
			s = s[:i-len(tempNum)] + s[i:]
			i -= len(tempNum)
			repeat, _ := strconv.Atoi(tempNum)
			repeatStk = append(repeatStk, repeat)
			leftIndexes = append(leftIndexes, i)
			tempNum = ""
			continue
		}
		if s[i] == ']' {
			// 弹出左括号的位置
			leftIndex := leftIndexes[len(leftIndexes)-1]
			leftIndexes = leftIndexes[:len(leftIndexes)-1]
			// 替换字符串
			validStr := s[leftIndex+1 : i]
			// 弹出重复的次数
			repeat := repeatStk[len(repeatStk)-1]
			repeatStk = repeatStk[:len(repeatStk)-1]
			decodeStr := ""
			for j := 0; j < repeat; j++ {
				decodeStr += validStr
			}
			s = s[:leftIndex] + decodeStr + s[i+1:]
			i = leftIndex + len(decodeStr) - 1
		}
	}
	return s
}

func main() {
	data := "100[a]"
	fmt.Println(decodeString(data))
}
