package main

import (
	"fmt"
)

// 去除目标值重复
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i + 1:]...)
			i--
		}
	}

	return len(nums)
}

func main() {
	res := []int{0, 0, 1, 1, 1, 2, 2, 4, 4, 2}
	dd := removeElement(res, 2)
	fmt.Println(dd)
}
