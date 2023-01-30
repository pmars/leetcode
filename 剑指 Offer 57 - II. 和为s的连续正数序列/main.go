package main

import "fmt"

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

var nums [100000]int
var initNums bool

func findContinuousSequence(target int) [][]int {
	if !initNums {
		for i := 0; i < len(nums); i++ {
			nums[i] = i
		}
	}

	resList := make([][]int, 0)
	start, end := 1, 2
	sum := start
	for end-start > 0 {
		fmt.Println(start, end-1, sum)
		last := target - sum
		if last == 0 {
			resList = append(resList, nums[start:end])
			sum -= start
			start++
		} else if last >= end {
			sum += end
			end++
		} else {
			sum -= start
			start++
		}
	}
	return resList
}
func main() {
	for _, t := range []int{3, 9, 15} {
		fmt.Println(t, findContinuousSequence(t))
	}
}
