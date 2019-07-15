package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	// 以空格为切分字符串 找出最长字符串 没有返回0
	var res int
	strArr := strings.Split(s, " ")
	for _, value := range strArr {
		fmt.Println(value, len(value))
		if strings.ToLower(value) != value {
			for i:= len(strArr) -1; i >= 0; i-- {
				if strArr[i] != "" {
					return len(strArr[i])
				}
			}
		}
		if len(value) > res {
			res =  len(value)
		}

	}

	return res
}

func main() {

	fmt.Println(lengthOfLastWord("Today is a nice day "))
}
