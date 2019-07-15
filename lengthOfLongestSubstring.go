package main

import (
	"strings"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var (
		tempStr string
		length int
	)

	for i := 0; i < len(s); i++ {
		if strings.Contains(tempStr, string(s[i])) {
			if len(tempStr) > length {
				length = len(tempStr)
			}
			i = i - len(tempStr) + 1
			tempStr = ""
		}
		tempStr += string(s[i])
	}

	if len(tempStr) > length {
		length = len(tempStr)
	}
	return length
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
