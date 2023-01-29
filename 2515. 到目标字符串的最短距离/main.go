package main

import "fmt"

// https://leetcode.cn/problems/shortest-distance-to-target-string-in-a-circular-array/

/*
分析：
	如果不存在则返回-1，查看是否存在，最简单的方法就是map
	而，每次移动只能往前或者往后一个距离，那么最短路径，就是index的差了
	可以通过map记录index，之后计算向前和向后的index差，取最小的就可以了
*/
func closetTarget(words []string, target string, startIndex int) int {
	wordIndexMap := make(map[string][]int, len(words)) // []int target在words里面可能有多个
	for i, w := range words {
		if _, ok := wordIndexMap[w]; !ok {
			wordIndexMap[w] = make([]int, 0)
		}
		wordIndexMap[w] = append(wordIndexMap[w], i)
	}

	targetIndexList, ok := wordIndexMap[target]
	// 如果target不再列表中，直接返回失败
	if !ok {
		return -1
	}

	minStep := 1000
	for _, targetIndex := range targetIndexList {
		start := startIndex
		if start > targetIndex {
			start, targetIndex = targetIndex, start
		}
		minStep1 := targetIndex - start
		minStep2 := len(words) - targetIndex + start
		minStep = min(minStep, min(minStep1, minStep2))
	}
	return minStep
}

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func main() {
	fmt.Println(closetTarget([]string{"hello", "i", "am", "leetcode", "hello"}, "hello", 1)) // 1
	fmt.Println(closetTarget([]string{"a", "b", "leetcode"}, "leetcode", 0))                 // 1
	fmt.Println(closetTarget([]string{"i", "eat", "leetcode"}, "aet", 0))                    // -1
}
