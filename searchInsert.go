package main

// 顺序插入与查找
func searchInsert(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}else if nums[i] > target {
			return i
		}

	}

	return i
}

func main() {

}