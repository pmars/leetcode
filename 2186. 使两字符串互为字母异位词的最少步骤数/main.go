package main

import "fmt"

// https://leetcode.cn/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/

/*
分析：
	既然字母是一样的，但顺序可以不一样，那么直接找到里面不同的字母就可以了
	初步的想法是给两个字符串直接按照字母排序，之后类似归并的方法，去除一样的，留下的不一样的字符数量就是结果了
	之后，看题目，说只有小写字母，那么直接做一个26大小的map来计算就可以了
*/
func minSteps(s string, t string) int {
	sCharMap := make(map[int32]int, 26)
	tCharMap := make(map[int32]int, 26)
	for _, c := range s {
		sCharMap[c]++
	}
	for _, c := range t {
		tCharMap[c]++
	}

	diff := 0
	for c := 'a'; c <= 'z'; c++ {
		diff += diffFunc(sCharMap[c], tCharMap[c])
	}
	return diff
}

func diffFunc(m, n int) int {
	if m > n {
		return m - n
	}
	return n - m
}

func main() {
	fmt.Println(minSteps("leetcode", "coats")) // 7
	fmt.Println(minSteps("night", "thing"))    // 0
}
