package main

import (
	"math/big"
	"fmt"
)

// -123  -> -321
func reverse(x int) int {
	xStr := fmt.Sprintf("%d", x)
	var res string
	end := 0
	if string(xStr[0]) == "-" {
		end = 1
		res += "-"
	}

	for i := len(xStr) - 1; i >= end; i-- {
		res += string(xStr[i])
	}

	num, flag := new(big.Int).SetString(res, 10)
	if !flag {
		return -1
	}

	if int(num.Int64()) < -2147483648 || int(num.Int64()) > 2147483647 {
		return 0
	}

	return int(num.Int64())
}


func main() {
	reverse(1534236469)
}