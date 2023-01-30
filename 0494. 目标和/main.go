package main

import "fmt"

//https://leetcode.cn/problems/target-sum/

// 最简单的方法，直接递归，但这种方法比较费时间，没有剪枝等等
func findTargetSumWaysDigui(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	return findTargetSumWaysDigui(nums[1:], target-nums[0]) + findTargetSumWaysDigui(nums[1:], target+nums[0])
}

/*
 * sum(P) 前面符号为+的集合；sum(N) 前面符号为减号的集合
 * 所以题目可以转化为
 * sum(P) - sum(N) = target
 * => sum(nums) + sum(P) - sum(N) = target + sum(nums)
 * => 2 * sum(P) = target + sum(nums)
 * => sum(P) = (target + sum(nums)) / 2
 * 因此题目转化为01背包，也就是能组合成容量为sum(P)的方式有多少种
 */
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum < target || sum < -target || (sum+target)%2 == 1 {
		return 0
	}
	w := (sum + target) / 2
	dp := make([]int, w+1)
	dp[0] = 1
	for _, num := range nums {
		for j := w; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[w]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{100}, -200))
}
