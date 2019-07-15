package main

import (
	"fmt"
	"strconv"
	"math/big"
)

func addToArrayForm(A []int, K int) []int {
	if len(A) < 0 {
		return nil
	}
	var str string
	for i := 0; i <len(A); i++ {
		str += fmt.Sprintf("%d", A[i])
	}

	num, _ := new(big.Int).SetString(str, 10)
	strB := new(big.Int).Add(num, big.NewInt(int64(K))).String()

	var res = make([]int, len(strB))
	for i := 0; i < len(strB); i++ {
		c, _ := strconv.ParseInt(string(strB[i]), 10, 64)
		res[i] = int(c)
	}

	fmt.Println(res)
	return res
}

func main() {

	addToArrayForm([]int{1,2,6,3,0,7,1,7,1,9,7,5,6,6,4,4,0,0,6,3},516)
}
