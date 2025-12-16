package main

import "fmt"

// 合并区间
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 0 || n == 1 {
		return intervals
	}

	// 插入排序
	for i := range intervals {
		temp := intervals[i]
		j := i
		for ; j > 0 && temp[0] < intervals[j-1][0]; j-- {
			intervals[j] = intervals[j-1]
		}
		intervals[j] = temp
	}

	result := make([][]int, 0)
	for k := range intervals {
		rLen := len(result)
		if rLen == 0 || intervals[k][0] > result[rLen-1][1] {
			result = append(result, intervals[k])
			continue
		}
		if result[rLen-1][1] < intervals[k][1] {
			result[rLen-1][1] = intervals[k][1]
		}
	}

	return result
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
