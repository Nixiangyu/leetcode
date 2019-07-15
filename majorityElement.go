package main

func majorityElement(nums []int) int {
	var temp = make(map[int]int)
	var t = len(nums) / 2
	var res int
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
		if temp[nums[i]] > t {
			res = nums[i]
		}
	}

	return res
}

func main() {
}
