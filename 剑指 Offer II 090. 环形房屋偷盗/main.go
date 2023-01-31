package main

import "fmt"

// https://leetcode.cn/problems/PzWKhm/

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(robInner(nums[1:]), robInner(nums[:len(nums)-1]))
}

func robInner(nums []int) int {
	res := make([]int, len(nums)+1)
	res[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		res[i] = max(res[i-1], res[i-2]+nums[i-1])
	}
	// fmt.Println(res)
	return res[len(nums)]
}

func main() {
	for _, n := range [][]int{
		{2, 3, 3},
		{2, 3, 3, 3},
		{2, 3, 3, 4, 10},
		{3, 3, 4, 10, 2, 2},
		{2, 3},
		{2},
		{200, 3, 140, 20, 10},
	} {
		fmt.Println(n, rob(n))
	}
}
