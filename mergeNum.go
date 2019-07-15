package main

import "sort"

func mergeNum(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i:=0
	for j := 0; j < len(nums1); j ++ {
		if nums1[j] == 0 && i < len(nums2){
			nums1[j] = nums2[i]
			i++
		}
	}

	sort.Ints(nums1)
}
func main() {
}
