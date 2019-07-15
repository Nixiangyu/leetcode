package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	result := make([][]int, 0, 0)
	indexs := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i < 4 {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}

	newEle := make([]int, len(indexs))
	copy(newEle, indexs)
	result = append(result, newEle)
	for {
		find := false
		for i := 0; i < len(nums)-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true
				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					var leftNums []int
					leftNums = indexs[:i]
					sum := 0
					for i := 0; i < len(leftNums); i++ {
						if leftNums[i] == 1 {
							sum++
						}
					}
					//将前sum个改为1，之后的改为0
					for i := 0; i < len(leftNums); i++ {
						if i < sum {
							leftNums[i] = 1
						} else {
							leftNums[i] = 0
						}
					}
				}
				newEle1 := make([]int, len(indexs))
				copy(newEle1, indexs)
				result = append(result, newEle1)
				break
			}
		}
		if !find {
			break
		}
	}
	fmt.Println(result)
	var t = make(map[string]int, 0)
	temp := make([][]int, 0)
	for _, v := range result {
		line := make([]int, 0)
		var t1 string
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}

		}
		sort.Ints(line)
		t1 = fmt.Sprintf("%d%d%d%d", line[0], line[1], line[2], line[3])
		if t[t1] == 0 && line[0]+line[1]+line[2]+line[3] == target {
			temp = append(temp, line)
			t[t1] = 1
		}
	}

	return temp
}

func main() {
	nums := []int{-500, -481, -480, -469, -437, -423, -408, -403, -397, -381, -379, -377,
		-353, -347, -337, -327, -313, -307, -299, -278, -265, -258, -235, -227, -225, -193, -192,
		-177, -176, -173, -170, -164, -162, -157, -147, -118, -115, -83, -64, -46, -36, -35, -11, 0,
		0, 33, 40, 51, 54, 74, 93, 101, 104, 105, 112, 112, 116, 129, 133, 146, 152, 157, 158, 166, 177,
		183, 186, 220, 263, 273, 320, 328, 332, 356, 357, 363, 372, 397, 399, 420, 422, 429, 433, 451, 464, 484, 485, 498, 499}

	fourSum(nums, 2139)
}
