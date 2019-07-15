package main

import (
	"fmt"
	//"strconv"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var str string
	var j, i = len(num2)-1, len(num1)-1
	for ; i >= 0; i-- {
		if j >= 0 {
		}
	}

	if j < 0 {
		if i > 0 {
			str = num1[:i+1] + str
		}
	}

	if i < 0 {
		if j > 0 {
			str = num2[:j+1] + str
		}
	}
	fmt.Println()
	return str
}

func main() {
	fmt.Println(string(int('*')+1), string(int('1')), int('9'))
	a,_ := strconv.ParseInt("1", 10, 64)
	fmt.Println("==", a)
	addStrings("1", "1")
}
