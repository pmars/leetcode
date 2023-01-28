package main

import "fmt"

/*
分析：
	这个题，上下左右，一个单位一秒，斜着四个方向也是一秒，那么我从(0,0)到一个(m,n)的时间不就是max(m,n)
	所以，这个题直接计算两点之间的差值，累加起来就可以了
*/

func max(n, m int) int {
	if n < 0 {
		n = -n
	}
	if m < 0 {
		m = -m
	}
	if n > m {
		return n
	}
	return m
}

func minTimeToVisitAllPoints(points [][]int) int {
	secs := 0
	for i := 1; i < len(points); i++ {
		xDiff := points[i][0] - points[i-1][0]
		yDiff := points[i][1] - points[i-1][1]
		secs += max(xDiff, yDiff)
	}
	return secs
}

func main() {
	for _, s := range [][][]int{
		{{1, 1}, {3, 4}, {-1, 0}},
	} {
		fmt.Println(s, minTimeToVisitAllPoints(s))
	}
}
