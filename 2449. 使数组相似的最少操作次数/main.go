package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-operations-to-make-arrays-similar/

/*
分析：
	将数字变大变小，即可和目标一致，最少的次数，那么肯定是最小的数变为最小的数
	而且，奇数肯定变的结果也是奇数，所以，这块，给奇偶分开来计算也可以
*/
func makeSimilar(nums []int, target []int) int64 {
	// 将奇偶分开
	numsJi := make([]int, 0)
	numsOu := make([]int, 0)
	targetJi := make([]int, 0)
	targetOu := make([]int, 0)

	for _, n := range nums {
		if n%2 == 0 {
			numsOu = append(numsOu, n)
		} else {
			numsJi = append(numsJi, n)
		}
	}
	for _, t := range target {
		if t%2 == 0 {
			targetOu = append(targetOu, t)
		} else {
			targetJi = append(targetJi, t)
		}
	}

	sort.Slice(numsJi, func(i, j int) bool {
		return numsJi[i] < numsJi[j]
	})
	sort.Slice(numsOu, func(i, j int) bool {
		return numsOu[i] < numsOu[j]
	})
	sort.Slice(targetJi, func(i, j int) bool {
		return targetJi[i] < targetJi[j]
	})
	sort.Slice(targetOu, func(i, j int) bool {
		return targetOu[i] < targetOu[j]
	})

	times := makeSimilarPart(numsJi, targetJi)
	times += makeSimilarPart(numsOu, targetOu)

	return times
}

// 奇数，偶数单独计算的方法
func makeSimilarPart(nums, target []int) int64 {
	var times int64
	for i := 0; i < len(nums); i++ {
		if nums[i] < target[i] {
			times += int64(target[i]-nums[i]) / 2
		}
	}
	return times
}

func main() {
	fmt.Println(makeSimilar([]int{0, 2, 6}, []int{2, 2, 4}))             // 1
	fmt.Println(makeSimilar([]int{8, 12, 6}, []int{2, 14, 10}))          // 2
	fmt.Println(makeSimilar([]int{1, 2, 5}, []int{4, 1, 3}))             // 1
	fmt.Println(makeSimilar([]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1})) // 0
}
