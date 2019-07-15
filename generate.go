package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	var res [][]int
	var t1 = []int{1}
	res = append(res, t1)
	if numRows == 1 {
		return res
	}

	var t2 = []int{1, 1}
	res = append(res, t2)
	if numRows == 2 {
		return res
	}

	// 超过3行
	for i := 2; i < numRows; i++ {
		var temp [i+1]int
		for j := 0; j < i + 1; j++ {
			if j == 0 || j == i{
				temp[j] = 1
			}else{
				temp[j] = res[i -1][j] + res[i-1][j-1]
			}

		}
	}

	return res

}

func main() {
}
