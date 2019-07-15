package main

import (
	"fmt"
)

// '(', ')', '{', '}', '[' and ']'
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}

	var temp map[string]int
	temp = make(map[string]int)

	var rule map[string]int
	rule = make(map[string]int)

	for _, value := range s {
		if string(value) == "(" || string(value) == "{" || string(value) == "[" {
			if string(value) == "(" {
				rule[string(value)] = 1
			}
			if string(value) == "{" {
				rule[string(value)] = 2
			}
			if string(value) == "[" {
				rule[string(value)] = 3
			}
			temp[string(value)] = 1
		} else {
			if string(value) == ")" {
				rule[string(value)] = 1
			}
			if string(value) == "}" {
				rule[string(value)] = 2
			}
			if string(value) == "]" {
				rule[string(value)] = 3
			}
			temp[string(value)] = 0
		}

	}

	var tempNoDis string
	for _, value := range s {
		if temp[string(value)] == 1 {
			tempNoDis += string(value)
			continue
		} else {
			if tempNoDis != "" {
				if rule[string(value)] == rule[string(tempNoDis[len(tempNoDis) - 1])] {
					tempNoDis = tempNoDis[:len(tempNoDis) - 1]

				}else{
					break
				}
			}else{
				return false
			}
		}
	}

	if tempNoDis != "" {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("{{}}"))
}