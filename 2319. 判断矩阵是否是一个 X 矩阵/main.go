package main

//https://leetcode.cn/problems/check-if-matrix-is-x-matrix/

func min(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func checkXMatrix(grid [][]int) bool {
	size := len(grid) - 1
	for i := range grid {
		for j := range grid[i] {
			minI, minJ := min(i, size-i), min(j, size-j)
			if minI == minJ && grid[i][j] == 0 {
				return false
			}
			if minI != minJ && grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {

}
