package main

import (
	"fmt"
)

// 121 ture 122 false 222 true
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	xStr := fmt.Sprintf("%d", x)
	for i := 0; i < len(xStr) / 2; i++ {
		if xStr[i] != xStr[len(xStr) - i -1] {
			return false
		}
	}

	return true
}

func main() {
	res := isPalindrome(115)
	fmt.Println(res)
}