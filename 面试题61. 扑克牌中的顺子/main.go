package main

import "fmt"

// https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/

func isStraight(nums []int) bool {
	minNum, maxNum, zeroNum := 100, 0, 0
	numList := make([]int, 14)
	for _, n := range nums {
		if minNum > n && n != 0 {
			minNum = n
		}
		if maxNum < n {
			maxNum = n
		}
		if n == 0 {
			zeroNum++
		}
		numList[n]++
	}

	if maxNum-minNum >= 5 {
		return false
	}
	if minNum > 13 { // 都是0
		return true
	}
	for i := minNum; i <= maxNum; i++ {
		if numList[i] > 1 {
			return false
		}
	}
	return true
}

func main() {
	for _, s := range [][]int{
		{1, 2, 3, 4, 5},
		{0, 0, 0, 0, 0},
		{13, 13, 13, 13, 13},
		{13, 0, 0, 0, 0},
	} {
		fmt.Println(s, isStraight(s))
	}
}
