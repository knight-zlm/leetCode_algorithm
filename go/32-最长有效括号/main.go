package main

import "fmt"

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
示例 1:
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/
var maxLen = 0

func dp(s string, deep int) {
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			dp(s[i+1:], deep+1)
		} else {
			if deep > 0 && deep+1 > maxLen {
				maxLen = deep + 1
			}
		}
	}
	return
}

func longestValidParentheses2(s string) int {
	maxLen := 0
	dp := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		// 前面有没有左括号与之对应
		if s[i] == ')' && i-1 >= 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			before := 0
			//需要看左括号前面是否有匹配的括号对
			if i-dp[i-1]-2 >= 0 {
				before = dp[i-dp[i-1]-2]
			}
			dp[i] = 2 + dp[i-1] + before
			if maxLen < dp[i] {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}

func longestValidParentheses(s string) int {
	maxLen := 0
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// 前面有没有左括号与之对应
		if s[i] == ')' && i-1 >= 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			before := 0
			//需要看左括号前面是否有匹配的括号对
			if i-dp[i-1]-2 >= 0 {
				before = dp[i-dp[i-1]-2]
			}
			dp[i] = 2 + dp[i-1] + before
			if maxLen < dp[i] {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}

func main() {
	max := longestValidParentheses("()(())")
	fmt.Println(max)
	max = longestValidParentheses(")()())")
	fmt.Println(max)
}
