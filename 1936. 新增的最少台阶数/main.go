package main

import "fmt"

// https://leetcode.cn/problems/add-minimum-number-of-rungs/
/*
分析：
	这个题也算是简单的了，每次有台阶就走，不够了，增加一个最大步数的台阶就可以了
	当然，这个方法超时了
*/
func addRungsTimeOut(rungs []int, dist int) int {
	times, current := 0, 0
	for _, s := range rungs {
		for current+dist < s {
			current += dist
			times++
		}
		current = s
	}
	return times
}

func addRungs(rungs []int, dist int) int {
	times, current := 0, 0
	for _, s := range rungs {
		if current+dist < s {
			times += (s - 1 - current) / dist
		}
		current = s
	}
	return times
}

func main() {
	fmt.Println(addRungs([]int{1, 3, 5, 10}, 2)) // 2
	fmt.Println(addRungs([]int{3, 6, 8, 10}, 3)) // 0
	fmt.Println(addRungs([]int{3, 4, 6, 7}, 2))  // 1
	fmt.Println(addRungs([]int{5, 1000000000}, 1))
}
