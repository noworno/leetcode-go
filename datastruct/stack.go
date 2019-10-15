package main

import (
	"fmt"
)

func main() {
	s := new(Stack)
	s.Push(5)
	s.Push("Hello World")
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	fmt.Println(s.Size())
}

type Stack struct {
	arr   []interface{}
	index int
}

func (s *Stack) Pop() interface{} {
	if s.index > 0 {
		s.index--
		return s.arr[s.index]
	}
	return nil
}

func (s *Stack) Push(x interface{}) {
	if len(s.arr) <= s.index {
		s.arr = append(s.arr, x)
	} else {
		s.arr[s.index] = x
	}
	s.index++
}

func (s *Stack) Size() int {
	return s.index
}
