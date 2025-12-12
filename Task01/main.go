package main

import "fmt"

func SingleNumber(nums []int) int {
	countMap := make(map[int]int)
	for i := range nums {
		count, exist := countMap[nums[i]]
		if exist {
			countMap[nums[i]] = count + 1
		} else {
			countMap[nums[i]] = 1
		}
	}

	// 从map中找出只出现一次的数字
	for key, value := range countMap {
		if value == 1 {
			return key
		}
	}

	return 0
}

func main() {
	nums := [7]int{1, 2, 2, 3, 3, 4, 4}
	fmt.Println(SingleNumber(nums[0:]))
}
