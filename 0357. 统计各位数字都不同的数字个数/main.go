package main

import "fmt"

// https://leetcode.cn/problems/count-numbers-with-unique-digits/

/*
分析：
	这个题n的范围太小了 0 <= n <= 8 直接做个switch就可以了
	当然，结果是需要提前跑出来的
	题目说到：各位数字都不同的数字
	也就是说直接排列组合搞起了
*/

// checkLen 可选择的数字个数
func countNumbersWithUniqueDigitsAnswer(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	count := 9
	for i, j := 9, n-1; j > 0; i, j = i-1, j-1 {
		count *= i
	}
	return count + countNumbersWithUniqueDigitsAnswer(n-1)
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	if n == 2 {
		return 91
	}
	if n == 3 {
		return 739
	}
	if n == 4 {
		return 5275
	}
	if n == 5 {
		return 32491
	}
	if n == 6 {
		return 168571
	}
	if n == 7 {
		return 712891
	}
	if n == 8 {
		return 2345851
	}
	if n == 9 {
		return 5611771
	}
	return 0
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, countNumbersWithUniqueDigitsAnswer(i))
	}
}
