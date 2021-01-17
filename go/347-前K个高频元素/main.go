package main

import (
	"container/heap"
	"fmt"
)

/**
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]


提示：
你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。
*/
func topKFrequent(nums []int, k int) []int {
	countNum := make(map[int]int, len(nums))
	for _, v := range nums {
		countNum[v]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, val := range countNum {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
