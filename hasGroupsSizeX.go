package main

import (
	"fmt"
)

func hasGroupsSizeX(deck []int) bool {
	var temp = make(map[int]int, 0)
	for _, v := range deck {
		temp[v]++
	}
	var t int
	if len(temp) > 0 {
		t = temp[deck[0]]
	}

	for _, v := range temp {
		if v != t {
			v1 := v
			t1 := t
			var key = -1
			var tmp int
			for {
				tmp = (v1 % t1)
				if tmp > 0 {
					v1 = t1
					t1 = tmp
				} else {
					key = t1
					break
				}
			}

			if key == -1 || key == 1 {
				return false
			}
		}

		t = v
	}

	return true
}

func main() {
	//fmt.Println(gcdx(4, 6))
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
}
