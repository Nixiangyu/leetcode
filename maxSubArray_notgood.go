package main

import (
	"fmt"
)

// 子数组最大和
func maxSubArray(nums []int) int {

	var res = nums[0]
	for j := 0; j < len(nums); j++ {
		var temp int
		for i := j; i < len(nums); i++ {
			temp += nums[i]
			if i == 0 {
				if nums[i] > res {
					res = nums[i]
				}
			}
			if temp > res {
				res = temp
			}
		}
	}

	return res
}


func main() {
	res := []int{ -1}

	fmt.Println(maxSubArray(res))
}
