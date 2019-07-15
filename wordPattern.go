package main

import (
	"strings"
	"fmt"
)

func wordPattern(pattern string, str string) bool {
	strArr := strings.Fields(str)
	if len(strArr) != len(pattern) {
		return false
	}

	keyStr := make(map[string]string, 0)
	keyPatten := make(map[string]string, 0)
	for i := 0; i < len(pattern); i++ {
		if i == 0 {
			keyStr[ string(pattern[i])] = strArr[i]
			keyPatten[strArr[i]] = string(pattern[i])
			continue
		}

		if keyStr[string(pattern[i])] != strArr[i] || keyPatten[strArr[i]] != string(pattern[i]) {
			if keyStr[string(pattern[i])] == "" && keyPatten[strArr[i]] == "" {
				keyStr[ string(pattern[i])] = strArr[i]
				keyPatten[strArr[i]] = string(pattern[i])
				continue
			}
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}
