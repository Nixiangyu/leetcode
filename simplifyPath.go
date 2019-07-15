package main

import (
	"strings"
	"fmt"
)

func simplifyPath(path string) string {
	if path == "" {
		return path
	}
	strPath := path
	strPath = strings.Replace(strPath,"//", "/", -1)
	fmt.Println(strPath)
	index := strings.LastIndex(strPath, ".")
	if index > 0 {
		strPath = strPath[index +1:]
	}
	fmt.Println(path, strPath, len(strPath))
	if len(strPath) > 1 {
		if string(strPath[len(strPath) - 1]) == "/" {
			strPath = strPath[:len(strPath) -1]
		}
	}

	return strPath
}

func main() {
    fmt.Println(simplifyPath("/home//foo/"))
}
