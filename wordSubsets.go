package main

import (
	"fmt"
	"strings"
)

/*// 效率太地 不接受
func wordSubsets(A []string, B []string) []string {

	for i := 0; i < len(B); i++ {
		for j := 0; j < len(A); j++ {
			var isCon bool
			if strings.Contains(A[j], B[i]) {
				continue
			}

			if len(B[i]) > 1 {
				for k := 0; k < len(B[i]); k++ {
					countA := strings.Count(A[j], string(B[i][k]))
					if countA <= 0{
						isCon = true
						break
					}

					countB := strings.Count(B[i], string(B[i][k]))
					if countA < countB{
						isCon = true
						break
					}
				}

				if !isCon {
					continue
				}
			}

			A = append(A[:j], A[j+1:]...)
			j--
		}
	}

	fmt.Println(A)

	return A
}*/

/*// 效率太地 不接受
func wordSubsets(A []string, B []string) []string {
	if len(B) == 0 {
		return A
	}

	var res []string
	for i := 0; i < len(A); i++ {
		var isCon bool
		for j := 0; j < len(B[len(B)-1]); j++ {
			countA := strings.Count(A[i], string(B[len(B)-1][j]))
			if countA <= 0{
				isCon = true
				break
			}

			countB := strings.Count(B[len(B)-1], string(B[len(B)-1][j]))
			if countA < countB{
				isCon = true
				break
			}
		}

		if isCon {
			continue
		}
		res = append(res, A[i])
	}

	return wordSubsets(A, B[:len(B)-1])
}*/

// 效率太地 不接受
func wordSubsets(A []string, B []string) []string {




	return wordSubsets(A, B)
}

//["amazon","apple","facebook","google","leetcode"]
//["lo","eo"]

func main() {
	fmt.Println("==", wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lo", "oo"}))
}
