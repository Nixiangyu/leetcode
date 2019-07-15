package main

import (
	"strings"
	"fmt"
	"regexp"
)

func backspaceCompare(S string, T string) bool {
	if (len(S) > 200 || len(S) < 1) && (len(T) > 200 || len(T) < 1) {
		return true
	}
	reg := regexp.MustCompile("^[A-Za-z]+$")
	strA := strings.Replace(S, "#", "", -1)
	strB := strings.Replace(T, "#", "", -1)
	if !reg.MatchString(strA) && !reg.MatchString(strB) {
		return true
	}

	if !reg.MatchString(strA) && reg.MatchString(strB) {
		return false
	}
	if reg.MatchString(strA) && !reg.MatchString(strB) {
		return false
	}
	if len(strA) == 0 && len(strB) == 0 {
		return true
	}
	fmt.Println(strA, strB)
	for i := 0; i < len(S); i++ {
		index := strings.Index(T, string(S[i]))
		if index < 0 {
			continue
		}
		if index+1 < len(T) {
			strS := S[i+1:]
			strT := T[index+1:]
			for j := 0; j < len(strS); i++ {
				if strings.Contains(strT, string(strS[i])) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	//^[A-Za-z]+$
	text := `Hello123`

	fmt.Println()
	//	fmt.Println(backspaceCompare("12", "112"))
}
