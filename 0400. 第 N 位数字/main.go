package main

import "fmt"

// https://leetcode.cn/problems/nth-digit/
func findNthDigit(n int) int {
	// start, end, nums
	// 1, 9, 9*1
	// 10, 99, 90*2
	// 100, 999, 900*3
	// 1000, 9999, 9000*4 ...

	start := 1
	for i := 1; ; i++ {
		start *= 10
		startNum, nums := start/10, 9*start/10*i
		if n > nums {
			n -= nums
			continue
		}

		num, res := startNum+(n-1)/i, 0
		// fmt.Println(num, n, i)
		for numIndex := (n - 1) % i; numIndex >= 0; numIndex-- {
			res = num / startNum
			num %= startNum
			startNum /= 10
		}
		return res
	}
}

func main() {
	for _, n := range []int{11, 10, 3} {
		fmt.Println(n, findNthDigit(n))
	}
}
