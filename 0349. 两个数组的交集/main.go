package main

import "fmt"

// https://leetcode.cn/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool)
	for _, n := range nums1 {
		numMap[n] = true
	}
	resList := make([]int, 0)
	for _, n := range nums2 {
		if numMap[n] {
			resList = append(resList, n)
			numMap[n] = false
		}
	}
	return resList
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
