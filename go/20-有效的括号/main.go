package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 1:
输入: "()"
输出: true
示例 2:
输入: "()[]{}"
输出: true
示例 3:
输入: "(]"
输出: false
示例 4:
输入: "([)]"
输出: false
示例 5:
输入: "{[]}"
输出: true
*/

var pair = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0, len(s))
	for _, val := range []byte(s) {
		_, ok := pair[val]
		if ok {
			stack = append(stack, val)
			continue
		}
		// 看最后一个是不是匹配的
		if len(stack) == 0 {
			return false
		}
		last := stack[len(stack)-1]
		must, _ := pair[last]
		if must != val {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("){"))
}
