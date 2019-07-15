package main

import (
	"fmt"
	"sort"
)

// 将以排序存在0的和以排序的2合并为1
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) < len(nums2) || len(nums1) < n || len(nums2) < n || len(nums1) < m || len(nums1) < (m+n) {

		return
	}
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}

	sort.Ints(nums1[:m+n])


	fmt.Println(nums1)

}

func main() {
	merge([]int{0, 0, 4, 0, 0, 0}, 3, []int{2, 6, 7}, 3)
}
