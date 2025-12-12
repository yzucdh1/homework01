package main

import (
	"errors"
	"fmt"
)

// 栈结构体
type Stack struct {
	data []rune
}

// 压栈
func (s *Stack) Push(val rune) {
	s.data = append(s.data, val)
}

// 出栈
func (s *Stack) Pop() (rune, error) {
	sLen := len(s.data)
	if sLen == 0 {
		return 0, errors.New("栈为空,无法出栈")
	}

	ret := s.data[sLen-1]
	s.data = s.data[:sLen-1]
	return ret, nil
}

// 有效的括号
func isValid(s string) bool {
	stack := &Stack{}
	arr := []rune(s)

	for i := range arr {
		char := arr[i]
		if char == '{' || char == '(' || char == '[' {
			stack.Push(char)
			continue
		}
		val, error := stack.Pop()
		if error != nil {
			return false
		}
		if char == '}' && val != '{' {
			return false
		} else if char == ')' && val != '(' {
			return false
		} else if char == ']' && val != '[' {
			return false
		}
	}

	return len(stack.data) == 0
}

func main() {
	s := "({)}"
	fmt.Println(isValid(s))
}
