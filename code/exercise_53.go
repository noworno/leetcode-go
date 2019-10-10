package main

import (
	"fmt"
)

func main() {
	arr1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr1))
}

func maxSubArray(nums []int) int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	maxres := nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = max(nums[i], res[i-1]+nums[i])
		if maxres < res[i] {
			maxres = res[i]
		}
	}
	//fmt.Println(res)
	return maxres
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
