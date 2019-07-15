package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return nil
	}

	var res [][]int
	var t1 = []int{1}
	res = append(res, t1)
	if rowIndex == 1 {
		return t1
	}

	var t2 = []int{1, 1}
	res = append(res, t2)
	if rowIndex == 2 {
		return t2
	}

	// 超过3行
	for i := 2; i < rowIndex; i++ {
		var temp [i+1]int
		for j := 0; j < i + 1; j++ {
			if j == 0 || j == i{
				temp[j] = 1
			}else{
				temp[j] = res[i -1][j] + res[i-1][j-1]
			}

		}
	}

	return res[rowIndex]
}

func main() {
}
