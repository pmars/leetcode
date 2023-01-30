package main

import "fmt"

// https://leetcode.cn/problems/4sum-ii/
/*
分析：
	想法1：
		最简单的想法，就是四层for，肯定能给结果弄出来， n最大是200，四层的话，16亿次，估计就timeout了
	想法2：
		优化一下，那么三层for，最后一层排好序，直接二分，复杂度 200*200*200*log200，也不是很理想
	想法3：
		前两层for，最后两个排好序，做两个游标，前后逼近，找对应的值，这个也就是 200*200*200*2，快了一些
	想法4：
		是不是，可以给两层的和放到map里面，之后两层for，这样复杂度，可以做到 200*200*log200，更快了
*/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	num12SumMap := make(map[int]int) // sum to count
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			num12SumMap[nums1[i]+nums2[j]]++
		}
	}
	count := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum := nums3[i] + nums4[j]
			if c, ok := num12SumMap[-sum]; ok {
				count += c
			}
		}
	}
	return count
}

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	fmt.Println(fourSumCount([]int{1}, []int{-2}, []int{-1}, []int{2}))
}
