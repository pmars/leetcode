package main

import "fmt"

// https://leetcode.cn/problems/most-visited-sector-in-a-circular-track/

/*
分析：
	中间转圈的都可以不计算，那么也就是开始和结束的地方是最多的
	最终结果是按照编号生序，注意一下从大号到小号的时候，小号的数字插入到前面就可以了
*/

func mostVisited(n int, rounds []int) []int {
	start, end := rounds[0], rounds[len(rounds)-1]

	// 分两种情况
	if start <= end {
		resList := make([]int, end-start+1)
		for i := 0; i <= end-start; i++ {
			resList[i] = start + i
		}
		return resList
	}
	// 从大数到小数，先给小数弄进来
	resList := make([]int, 0)
	for i := 0; i < end; i++ {
		resList = append(resList, i+1)
	}
	for i := start; i <= n; i++ {
		resList = append(resList, i)
	}
	return resList
}

func main() {
	fmt.Println(mostVisited(4, []int{1, 3, 1, 2}))
	fmt.Println(mostVisited(7, []int{1, 3, 5, 7}))
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
	fmt.Println(mostVisited(6, []int{5, 1, 2, 1, 2, 1, 2, 1, 2}))
}
