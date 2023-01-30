package main

import "fmt"

// https://leetcode.cn/problems/convert-1d-array-into-2d-array/

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := make([][]int, m)
	for i, j := 0, 0; i < m; i, j = i+1, j+1 {
		println(i, j)
		res[j] = original[i*n : i*n+n]
	}
	return res
}

func main() {
	fmt.Println(construct2DArray([]int{1, 2, 3}, 1, 3))
	fmt.Println(construct2DArray([]int{1, 2}, 1, 1))
	fmt.Println(construct2DArray([]int{3}, 1, 2))
	fmt.Println(construct2DArray([]int{1, 2, 3, 4}, 2, 2))
	fmt.Println(construct2DArray([]int{1, 2, 3, 4, 5, 6}, 3, 2))
}
