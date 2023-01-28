package main

import "fmt"

// https://leetcode.cn/problems/minimum-moves-to-move-a-box-to-their-target-location/

/*
分析：
	这个题，虽然标的是困难的，但应该就是一个广搜，之后设计一个m*n的数组，记录最小步数，进行剪枝和确定结果
	在计算过程中，需要注意的是，push的方向判断
*/
func minPushBox(grid [][]byte) int {
	resMap := make([][]int, len(grid))
	direMap := make([][][]bool, len(grid)) // 某一点推动的方向
	var startX, startY int                 // 箱子的位置
	var endX, endY int                     // 终点的位置
	var personX, personY int               // 人的位置
	for i, list := range grid {
		resMap[i] = make([]int, len(list))
		direMap[i] = make([][]bool, len(list))
		for j, c := range list {
			direMap[i][j] = make([]bool, 4)
			switch c {
			case 'T':
				endX, endY = i, j
			case 'S':
				personX, personY = i, j
			case 'B':
				startX, startY = i, j
			}
		}
	}
	// fmt.Println(startX, startY, endX, endY, personX, personY)

	var directions = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	nextPoints := [][]int{{startX, startY, personX, personY}}
	for len(nextPoints) > 0 {
		startX, startY, personX, personY = nextPoints[0][0], nextPoints[0][1], nextPoints[0][2], nextPoints[0][3]
		nextPoints = nextPoints[1:]

		for di, d := range directions {
			if direMap[startX][startY][di] {
				continue
			}

			nextX, nextY := startX+d[0], startY+d[1]             // 箱子推动的下一个位置
			personNextX, personNextY := startX-d[0], startY-d[1] // 人站在哪里推动这个箱子
			if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
				continue
			}
			if personNextX < 0 || personNextY < 0 || personNextX >= len(grid) || personNextY >= len(grid[0]) {
				continue
			}
			// 如果这两个位置，任何一个是墙，就不用推了
			if grid[nextX][nextY] == '#' || grid[personNextX][personNextY] == '#' {
				continue
			}
			// 如果推动的下一个位置，之前有步数，并且是比较小的，那么也不用往那边推了
			//if resMap[nextX][nextY] != 0 && resMap[nextX][nextY] <= resMap[startX][startY]+1 {
			//	continue
			//}
			// 如果人不可以到推箱子的位置，也不用计算了
			if !personMove(grid, startX, startY, personX, personY, personNextX, personNextY) {
				continue
			}

			// 限制条件没有问题，则下一个位置可以
			// 先看看是不是有结果了
			if nextX == endX && nextY == endY {
				return resMap[startX][startY] + 1
			}

			direMap[startX][startY][di] = true
			resMap[nextX][nextY] = resMap[startX][startY] + 1
			nextPoints = append(nextPoints, []int{nextX, nextY, personNextX, personNextY})
		}
	}

	return -1
}

// person从(personX, personY) 到 (personNextX, personNextY) 是否可以过去
func personMove(grid [][]byte, boxX, boxY, personX, personY, personNextX, personNextY int) bool {
	if personX == personNextX && personY == personNextY {
		return true
	}

	var directions = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	pointList := []int{personX*100 + personY}
	targetMap := map[int]bool{personX*100 + personY: true}
	for len(pointList) > 0 {
		personX, personY = pointList[0]/100, pointList[0]%100
		pointList = pointList[1:]
		for _, d := range directions {
			nextX, nextY := personX+d[0], personY+d[1]
			if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
				continue
			}
			if nextX == boxX && nextY == boxY {
				continue
			}
			if grid[nextX][nextY] == '#' {
				continue
			}

			if nextX == personNextX && nextY == personNextY {
				return true
			}
			pointInt := nextX*100 + nextY
			if targetMap[pointInt] {
				continue
			}
			targetMap[pointInt] = true
			pointList = append(pointList, nextX*100+nextY)
		}
	}
	return false
}

func main() {
	fmt.Println(minPushBox([][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '.', '.', '#', '#'},
		{'#', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}))

	fmt.Println(minPushBox([][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '#', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'}}))

	fmt.Println(minPushBox([][]byte{
		{'#', '.', '.', '#', '#', '#', '#', '#'},
		{'#', '.', '.', 'T', '#', '.', '.', '#'},
		{'#', '.', '.', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', '#', '.', 'S', '#'},
		{'#', '.', '.', '#', '#', '#', '#', '#'}}))
	fmt.Println(minPushBox([][]byte{
		{'#', '.', '.', '#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', '#', '.', '#', '.', '.', '#'},
		{'#', '.', '.', '#', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', '.', '#', '.', 'S', '#'},
		{'#', '.', '.', '#', '.', '#', '#', '#', '#'}}))

	fmt.Println(minPushBox([][]byte{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '#', '#', '#', '#'},
		{'#', '.', '#', '#', '#', '#', '.', '#', '#', '#', '#', '.', '#', '#', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '#', 'T', '#', '.', '.', '#', '#', '#', '.'},
		{'#', '.', '.', '.', '#', '.', '.', '.', '.', '.', '.', '.', '#', '#', '#', '.'},
		{'#', '.', '.', '.', '.', '.', 'B', '.', '.', '.', '.', '.', '#', '#', '#', '.'},
		{'#', '.', '#', '#', '#', '#', '#', '#', '#', '#', '#', '.', '#', '#', '#', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', 'S', '.', '.', '.', '.', '.', '.', '.'},
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'}}))

	fmt.Println(minPushBox([][]byte{
		{'#', '#', '#', '#', '#', '#', '#'},
		{'#', 'S', '#', '.', 'B', 'T', '#'},
		{'#', '#', '#', '#', '#', '#', '#'}}))

	fmt.Println(minPushBox([][]byte{{'S'}, {'B'}, {'T'}, {'.'}, {'#'}}))

}
