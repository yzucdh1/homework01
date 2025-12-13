package main

import "fmt"

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	i := 0
	for j := 1; j < n; j++ {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
	}

	return i + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
