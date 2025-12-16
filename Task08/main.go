package main

import "fmt"

// 两数之和
func twoSum(nums []int, target int) []int {
	tempMp := make(map[int]int)
	result := make([]int, 2)
	for i := range nums {
		v, exist := tempMp[target-nums[i]]
		if exist {
			result[0] = v
			result[1] = i
		} else {
			tempMp[nums[i]] = i
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
