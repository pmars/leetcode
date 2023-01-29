package main

import "fmt"

// https://leetcode.cn/problems/contains-duplicate/

/*
分析：
	数字的空间是在9次方以内，太大了，明显不能用数组直接弄，数字的数量在5次方，10w内，还可以 ，直接上map就可以了
	尝试直接生成5次方大小的map，试一试空间换时间的速度
*/
func containsDuplicate(nums []int) bool {
	// numMap := make(map[int]bool, 100000) // 空间换时间
	numMap := make(map[int]bool)
	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			return true
		}
		numMap[n] = true
	}
	return false
}

func main() {

	for _, s := range [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	} {
		fmt.Println(s, containsDuplicate(s))
	}
}
