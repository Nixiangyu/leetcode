package main

import (
	"strconv"
	"fmt"
	"math"
)

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	if len(digits) == 1 && digits[0] >= 10 {
		return nil
	}

	if digits[len(digits)-1]+1 < 10 {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
		return digits
	}

	var num int64
	for i := 0; i < len(digits); i++ {
		if digits[i] >= 10 && i != len(digits)-1 {
			return nil
		}
		num += int64(math.Pow(10, float64(len(digits)-i-1))) * int64(digits[i])

	}
	fmt.Println(num + 1)

	numStr := strconv.Itoa(int(num + 1))
	for i := 0; i < len(numStr); i++ {
		var temp, _ = strconv.ParseInt(string(numStr[i]), 10, 64)
		if i > len(digits)-1 {
			digits = append(digits, int(temp))
		}
		digits[i] = int(temp)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
