package main

import (
	"fmt"
)


/*func twoSum(nums []int, target int) []int {
	var res []int
	var temp  map[int]int
	temp = make(map[int]int)

	for key, value := range nums {
		temp[value] = key
	}

	for key, value := range nums {
		if temp[target - value] != 0 && temp[target - value] != key {
			res = append(res, key)
			res = append(res, temp[target - value])
			break
		}
	}

	return res
}*/

// 一组数组 给定一个目标值 找到和为目标值的两个下标
func twoSum(nums []int, target int) []int {
	var res []int
	var temp  map[int]int
	temp = make(map[int]int)

	for key, value := range nums {
		fmt.Println(value, key, temp[target - value])
		if (temp[target - value] != 0 && target - value == value) || (temp[target - value] != 0 && temp[target - value] - 1 != key) {
			res = append(res, key)
			res = append(res, temp[target - value] - 1)
			break
		}
		temp[value] = key + 1
	}
	return res
}

func main() {
	testArr := []int{3, 3}
	twoSum(testArr, 6)
}
