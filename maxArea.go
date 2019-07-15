package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var (
		temp int
	)
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
		 if  height[i] *j > temp {
			 temp =  height[i] *j
		 }

		}
	}

	fmt.Println(temp)
	return temp
}

func main() {
	// []
	var a = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(a)
}
