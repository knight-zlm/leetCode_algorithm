package main

import (
	"fmt"
)

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

//

func dp(data []int) []int {
	if len(data) == 1 {
		return data
	}
	mid := (0 + len(data)) / 2
	left := dp(data[0:mid])
	right := dp(data[mid:])
	//合并
	leftIndex := 0
	rightIndex := 0
	newList := make([]int, 0, len(left)+len(right))
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] > right[rightIndex] {
			newList = append(newList, right[rightIndex])
			rightIndex++
		} else {
			newList = append(newList, left[leftIndex])
			leftIndex++
		}
	}
	if leftIndex != len(left) {
		newList = append(newList, left[leftIndex:]...)
	} else {
		newList = append(newList, right[rightIndex:]...)
	}
	return newList
}

// 归并排序
func MergeSort(data []int) {

}

func swap(nums []int, left, right int) {
	temp := nums[left]
	nums[left] = nums[right]
	nums[right] = temp
}

func nextPermutation(nums []int) {
	temp := nums[len(nums)-1]
	changeIndex := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] >= temp {
			temp = nums[i]
			continue
		}
		temp = nums[i]
		changeIndex = i
		break
	}
	// 证明不是从大到小
	if changeIndex != -1 {
		nextIndex := len(nums)
		// 把找到的小于后面一个数的值和后面和它最接近的比它大一点的值交换位置
		for i := changeIndex + 1; i < len(nums); i++ {
			if nums[i] <= temp {
				nextIndex = i
				break
			}
		}
		// 交换位置
		swap(nums, changeIndex, nextIndex-1)
	}
	left := changeIndex + 1
	right := len(nums) - 1
	for left < right {
		swap(nums, left, right)
		left++
		right--
	}
}

func main() {
	//data2 := dp([]int{1, 3, 6, 10, 4, 2, 5, 8, 7})
	//fmt.Println(data2)
	data1 := []int{1, 5, 1}
	nextPermutation(data1)
	fmt.Println(data1)
}
