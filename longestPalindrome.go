package main

import (
	"fmt"
)

func isPalind(s string) bool {
	var (
		i = 0
		j = len(s) - 1
	)
	for i <= j {
		if string(s[i]) != string(s[j]) {
			return false
		}
		i++
		j--
	}

	return true
}

func longestPalindrome(s string) (res string) {
	if len(s) < 2 {
		return s
	}

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if len(res) > len(s[i:j+1]) {
				continue
			}
			if isPalind(s[i:j+1]) {
				res = s[i:j+1]
			}
		}
	}

	if res == "" {
		return string(s[len(s)-1])
	}
	return res
}

func main() {
	fmt.Println("===", longestPalindrome("ac"))
}
