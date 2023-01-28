package main

import "fmt"

// https://leetcode.cn/problems/running-sum-of-1d-array/submissions/
func runningSum(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}

func main() {
	for _, s := range [][]int{
		{1, 2, 3},
		{2, 1, 6, 4},
		{6, 1, 7, 4, 1},
	} {
		fmt.Println(s, runningSum(s))
	}
}
