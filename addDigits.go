package main

import "fmt"

func addDigits(num int) int {
	var temp int
	for {
		if num == 0 {
			if temp < 10 {
				return temp
			}else{
				num = temp
				temp = 0
			}
		}
		temp += num % 10
		num = num / 10

	}

	return temp
}

func main() {
	fmt.Println(addDigits(38))
}
