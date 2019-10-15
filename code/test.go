package main

import (
	"fmt"
)

func main() {
	fmt.Println(queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}, []int{0, 0}))
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	res := make([][]int, 0)
	i := king[0]
	j := king[1]
	for i >= 0 {
		if isAttack(queens, i-1, j) {
			r := []int{i - 1, j}
			res = append(res, r)
			break
		}
		i--
	}
	fmt.Println("a", res)
	i = king[0]
	j = king[1]
	for i < 8 {
		if isAttack(queens, i+1, j) {
			r := []int{i + 1, j}
			res = append(res, r)
			break
		}
		i++
	}
	fmt.Println("b", res)
	i = king[0]
	j = king[1]
	for j >= 0 {
		if isAttack(queens, i, j-1) {
			r := []int{i, j - 1}
			res = append(res, r)
			break
		}
		j--
	}
	fmt.Println("c", res)
	i = king[0]
	j = king[1]
	for j < 8 {
		if isAttack(queens, i, j+1) {
			r := []int{i, j + 1}
			res = append(res, r)
			break
		}
		j++
	}
	fmt.Println("d", res)
	i = king[0]
	j = king[1]
	for i >= 0 && j >= 0 {
		if isAttack(queens, i-1, j-1) {
			r := []int{i - 1, j - 1}
			res = append(res, r)
			break
		}
		i--
		j--
	}
	fmt.Println("e", res)
	i = king[0]
	j = king[1]
	for i < 8 && j < 8 {
		if isAttack(queens, i+1, j+1) {
			r := []int{i + 1, j + 1}
			res = append(res, r)
			break
		}
		i++
		j++
	}
	fmt.Println("f", res)
	i = king[0]
	j = king[1]
	for j >= 0 && i < 8 {
		if isAttack(queens, i+1, j-1) {
			r := []int{i + 1, j - 1}
			res = append(res, r)
			break
		}
		j--
		i++
	}
	fmt.Println("g", res)
	i = king[0]
	j = king[1]
	for i >= 0 && j < 8 {
		if isAttack(queens, i-1, j+1) {
			r := []int{i - 1, j + 1}
			res = append(res, r)
			break
		}
		i--
		j++
	}
	return res
}

func isAttack(queens [][]int, i, j int) bool {
	for _, v := range queens {
		if v[0] == i && v[1] == j {
			return true
		}
	}
	return false
}
