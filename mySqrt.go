package main

import "fmt"

/*
func mySqrt(x int) int {
	for i := 0; ; i++ {
		if i*i > x {
			return i - 1
		}
	}

	return -1
}
*/

func mySqrt(x int) int {
	for i := x / 2; x > 1 && i >= 0; i-- {
		if i * i <= x {
			fmt.Println("==>>",i)
			return i
		}

		fmt.Println("=ddd=>>",i)
	}

	return x
}

func main() {
	fmt.Println(mySqrt(44641))
}

