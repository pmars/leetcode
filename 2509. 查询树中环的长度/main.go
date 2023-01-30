package main

import "fmt"

// https://leetcode.cn/problems/cycle-length-queries-in-a-tree/

// 这种方法，给一条边先走到底，之后另一条边走，碰到了走过的边，直接累和就行，但时间有点多，有一些计算浪费了
func cycleLengthQueriesV1(n int, queries [][]int) []int {
	res := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		start, end := queries[i][0], queries[i][1]
		stepMap := make(map[int]int)
		if start > end {
			start, end = end, start
		}
		for step := 1; start > 0; step, start = step+1, start/2 {
			stepMap[start] = step
		}
		for step := 0; end > 0; step, end = step+1, end/2 {
			if stepMap[end] > 0 {
				res[i] = stepMap[end] + step
				break
			}
		}
	}
	return res
}

// 这个方法，就是，谁大谁走，直到遇上，去除了一些浪费的步伐
func cycleLengthQueries(n int, queries [][]int) []int {
	res := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		start, end := queries[i][0], queries[i][1]
		startStep, endStep := 0, 0

		for start != end {
			if start > end {
				startStep++
				start /= 2
			} else {
				endStep++
				end /= 2
			}
		}
		res[i] = 1 + startStep + endStep
	}
	return res
}

func main() {
	fmt.Println(cycleLengthQueries(3, [][]int{{5, 3}, {4, 7}, {2, 3}, {1, 1}, {5, 7}, {7, 7}}))
}
