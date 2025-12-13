package main

import (
	"fmt"
	"strings"
)

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	minLenStr := strs[0]
	for i := 1; i < len(strs); i++ {
		strLen := len(strs[i])
		if strLen < minLen {
			minLen = strLen
			minLenStr = strs[i]
		}
	}

	minLenRune := []rune(minLenStr)
	var builder strings.Builder
	for j := range minLenRune {
		char := minLenRune[j]
		toAdd := true
		for k := range strs {
			strRune := []rune(strs[k])
			if strRune[j] != char {
				toAdd = false
				break
			}
		}
		if !toAdd {
			break
		}
		builder.WriteRune(char)
	}

	return builder.String()
}

func main() {
	str := []string{"cir", "car"} //"dog","racecar","car"
	fmt.Println(longestCommonPrefix(str))
}
