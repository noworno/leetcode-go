package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	m := make(map[int]int)
	for i, v := range nums {
		av, ok := m[target-v]
		if ok {
			res[0] = av
			res[1] = i
			break
		}
		m[v] = i
	}
	//fmt.Println(m)
	return res
}
