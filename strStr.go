package main

import (
	"strings"
	"fmt"
)

func strStr(haystack string, needle string) int {
	index := strings.Index(haystack, needle)
	return index
}

func main() {
	fmt.Println(strStr())
}
