package main

import "fmt"

/*
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

示例 1：
输入：x = 1, y = 4
输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。

示例 2：
输入：x = 3, y = 1
输出：1
*/
func hammingDistance(x int, y int) int {
	req := 0
	for x != 0 || y != 0 {
		xh := 0
		if x != 0 {
			xh = x & 1
		}
		yh := 0
		if y != 0 {
			yh = y & 1
		}
		if xh != yh {
			req++
		}
		x = x >> 1
		y = y >> 1
	}

	return req
}

func main() {
	fmt.Println(hammingDistance(1, 3))
}
