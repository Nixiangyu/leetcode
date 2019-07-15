package main

import "fmt"

// PAYPALISHIRING
func convert(s string, numRows int) string {
	var(
		k = 0
		tempstr string
		res [][]string
	)
	for j:=0; k < len(s); j++{
		var tempStr = make([]string, numRows)
		if j %(numRows - 1) == 0 {
			for i:=numRows - 1; i >= 0 && k < len(s); i-- {
				tempStr[i] = string(s[k])
				fmt.Println(i)
				k++
			}

		}else{
			tempStr[j%(numRows-1)] = string(s[k])
			k++
		}
		res = append(res, tempStr)
	}

	for i:=numRows - 1; i >= 0; i--{
		for j:=0; j < len(res) ; j++ {
			fmt.Println(i, j)
			if res[j][i] != "" {
				tempstr += res[j][i]
			}
		}
	}

	return tempstr
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
