package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Println(longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

/*
  动态规划方法
*/
func longestSubsequence(arr []int, difference int) int {
	res := make([]int, len(arr))
	res[0] = 1
	var count int
	result := 0
	for i := 1; i < len(arr); i++ {
		count = 1
		for j := i - 1; j >= 0; j-- {
			if arr[i] == arr[j]+difference {
				count += res[j]
				break
			}
		}
		res[i] = count
		if result < count {
			result = count
		}
	}
	return result
}
