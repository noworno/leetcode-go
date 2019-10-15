package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
}

func uniquePaths(m int, n int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
		graph[i][0] = 1
		for j := 0; j < m; j++ {
			graph[0][j] = 1
			if i > 0 && j > 0 {
				graph[i][j] = graph[i-1][j] + graph[i][j-1]
			}
		}
	}
	return graph[n-1][m-1]
}
