package main

import (
	"fmt"
)

func main() {
	n := 5
	fmt.Println("Result --> ", climbStairs(n))
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	l := make([]int, n)
	l[0] = 1
	l[1] = 2
	for i := 2; i < n; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	return l[n-1]
}
