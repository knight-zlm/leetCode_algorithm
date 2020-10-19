package main

func dp(nums []int, left, right, total int) [][]int {
	out := make([][]int, 0, len(nums)/2)
	for left < right {
		dif := nums[left] + nums[right] - total
		if dif == 0 {
			out = append(out, []int{left, right})
		} else if dif > 0 {

		} else {

		}
		switch temp - total {
		case:

		}
		if nums[left]+nums[right] == total {
		}
	}
}

// 三数之和
func threeSum(nums []int) [][]int {
	for outIndex := range nums {
		num := nums[outIndex]
		for innerIndex := range nums[outIndex+1:] {

		}
	}
	return nil
}

func main() {

}
