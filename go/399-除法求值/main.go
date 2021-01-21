package main

import "fmt"

/**
给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

示例 1：
输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
条件：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]

示例 2：
输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]

示例 3：
输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出：[0.50000,2.00000,-1.00000,-1.00000]
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 需要一个哈希表做字符和编号的对应
	equationsMap := make(map[string]int, len(equations)*2)
	id := 0
	for _, v := range equations {
		val1 := v[0]
		if _, ok := equationsMap[val1]; !ok {
			equationsMap[val1] = id
			id++
		}
		val2 := v[1]
		if _, ok := equationsMap[val2]; !ok {
			equationsMap[val2] = id
			id++
		}
	}
	parent := make([]int, id)
	weight := make([]float64, id)
	for i := 0; i < id; i++ {
		parent[i] = i
		weight[i] = 1
	}
	// 查找函数
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			f := find(parent[x])
			weight[x] *= weight[parent[x]]
			parent[x] = f
		}
		return parent[x]
	}
	// 合并函数
	var merge func(int, int, float64)
	merge = func(from int, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		if fFrom != fTo {
			parent[fFrom] = fTo
			weight[fFrom] = val * weight[to] / weight[from]
		}
	}

	for i, v := range equations {
		merge(equationsMap[v[0]], equationsMap[v[1]], values[i])
	}
	ans := make([]float64, len(queries))
	for i, v := range queries {
		from, ok := equationsMap[v[0]]
		to, ok2 := equationsMap[v[1]]
		if !ok || !ok2 {
			ans[i] = -1.0
			continue
		}
		fFrom, fTo := find(from), find(to)
		if fFrom != fTo {
			ans[i] = -1.0
			continue
		}
		ans[i] = weight[from] / weight[to]
	}

	return ans
}

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(equations, values, queries))
}
