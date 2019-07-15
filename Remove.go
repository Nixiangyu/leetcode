package main

import "fmt"

// 去重复 在数组内
func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums) - 1; i++ {
		if i != len(nums) - 1 && nums[i] == nums[i + 1] {
			nums = append(nums[:i], nums[i + 1:]...)
			i--
		}
	}

	return len(nums)
}

func main() {
	res := []int{0, 0, 1, 1, 1, 2, 2, 4, 4, 5}
	dd := removeDuplicates(res)
	fmt.Println(dd)
}
