package main

import (
	"math"
	"fmt"
	"math/big"
	"strconv"
)


// 2进制转10进制
func twoBinary(s string) *big.Float  {
	var res1  = big.NewFloat(0)
	for i := len(s) -1; i >= 0; i-- {
		if string(s[i]) == "1" {
			temp := math.Pow(2, float64(len(s) - i - 1))
			res1.Add(res1, big.NewFloat(temp))
		}
	}

	return res1
}

// 10进制转2进制
func hhh(num *big.Float) string {
	numStr := num.String()

}
func addBinary(a string, b string) string {
	if a == "" || b == "" {
		return ""
	}
	// 判断字符串 将二进制变成10进制 然后再将10进制变成二进制 返回字符串
	ten := new(big.Float).Add(twoBinary(a), twoBinary(b))
	return tenBinary(ten)
}
// "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"
func main() {
	fmt.Println(twoBinary("1010"))
	fmt.Println(twoBinary("110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
	//fmt.Println(tenBinary("1010"))

	fmt.Println(addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101","110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}
