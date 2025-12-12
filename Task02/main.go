package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	// 将数字转成字符串
	nStr := strconv.Itoa(n)
	nArr := []rune(nStr)
	strLen := len(nArr)
	rArr := make([]rune, strLen)

	for i, j := strLen-1, 0; i >= 0; i-- {
		rArr[j] = nArr[i]
		j++
	}

	rStr := string(rArr)
	return nStr == rStr
}

func main() {
	n := 10
	fmt.Println(isPalindrome(n))
}
