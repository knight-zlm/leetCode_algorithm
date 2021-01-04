package main

import (
	"fmt"
	"strings"
)

/**
删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
说明: 输入可能包含了除 ( 和 ) 以外的字符。
示例 1:
输入: "()())()"
输出: ["()()()", "(())()"]

示例 2:
输入: "(a)())()"
输出: ["(a)()()", "(a())()"]

示例 3:
输入: ")("
输出: [""]
*/

func recurse(s string, index, leftCount, rightCount, leftRem, rightRem int, expr []string, result map[string]int) {
	if index == len(s) {
		if leftRem == 0 && rightRem == 0 {
			ans := strings.Join(expr, "")
			result[ans] = 1
		}
	} else {
		if s[index] == '(' && leftRem > 0 {
			recurse(s, index+1, leftCount, rightCount, leftRem-1, rightRem, expr, result)
		} else if s[index] == ')' && rightRem > 0 {
			recurse(s, index+1, leftCount, rightCount, leftRem, rightRem-1, expr, result)
		}
		expr = append(expr, string(s[index]))
		if s[index] != '(' && s[index] != ')' {
			recurse(s, index+1, leftCount, rightCount, leftRem, rightRem, expr, result)
		} else if s[index] == '(' {
			recurse(s, index+1, leftCount+1, rightCount, leftRem, rightRem, expr, result)
		} else if s[index] == ')' && leftCount > rightCount {
			recurse(s, index+1, leftCount, rightCount+1, leftRem, rightRem, expr, result)
		}
		expr = expr[:len(expr)-1]
	}
}

func removeInvalidParentheses(s string) []string {
	left := 0
	right := 0
	for _, char := range s {
		if char == '(' {
			left += 1
		} else if char == ')' {
			if left == 0 {
				right += 1
			}
			if left > 0 {
				left -= 1
			}
		}
	}
	result := make(map[string]int)
	expr := make([]string, 0, len(s))
	recurse(s, 0, 0, 0, left, right, expr, result)
	ans := make([]string, 0, len(s))
	for k := range result {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	data := "()())()"
	fmt.Println(removeInvalidParentheses(data))
}
