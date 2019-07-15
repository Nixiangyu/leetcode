package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var tmp = make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				tmp[nums[i]+nums[j]+nums[k]] = 1
			}
		}
	}

	var t float64
	var flag bool
	var t1 int
	for k, v := range tmp {
		if v == 0 {
			continue
		}

		if !flag {
			t =math.Abs(float64(k) - float64(target))
			t1 = k
			flag = true
			continue
		}

		if t > math.Abs(float64(k)-float64(target)) {
			t = math.Abs(float64(k) - float64(target))
			t1 = k
		}
	}

	return t1
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
