package main

import (
	"fmt"
	"strings"
)

// 提取共同字符串
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var temp string
	for i := 0; i < len(strs[0]); i++ {
		temp += string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if strings.Index(strs[j], temp) != strings.Index(strs[0], temp){
				return temp[:len(temp) - 1]
			}
		}
	}

	return temp
}

func main() {
	str := []string{"flower","flow","flight"}
	res := longestCommonPrefix(str)
	fmt.Println(res)
}