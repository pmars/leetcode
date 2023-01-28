package main

import "fmt"

// https://leetcode.cn/problems/route-between-nodes-lcci/

/*
分析：
	有向图，找连通性，不带权重，也算是简单了，golang的话，直接map做一个连通性的字典就可以了
	数据处理的话，可以先给graph，做成二维数组，方便后面通过start来构建这个map
	题目中有平行边，或者环，增加一个处理过的点的map就可以过滤了
*/
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 先处理graph，做成一个边的二维数组
	graphDataMap := make(map[int][]int)
	for _, g := range graph {
		if _, ok := graphDataMap[g[0]]; !ok {
			graphDataMap[g[0]] = make([]int, 0)
		}
		graphDataMap[g[0]] = append(graphDataMap[g[0]], g[1])
	}

	if _, ok := graphDataMap[start]; !ok {
		return false
	}

	targetList := graphDataMap[start]
	doneMap := map[int]bool{
		start: true,
	}

	for len(targetList) > 0 {
		first := targetList[0]
		targetList = targetList[1:]

		if doneMap[first] {
			continue
		}

		if first == target {
			return true
		}

		if list, ok := graphDataMap[first]; ok {
			for _, t := range list {
				if t == target {
					return true
				} else if doneMap[t] {
					continue
				} else {
					targetList = append(targetList, t)
				}
			}
		}

		doneMap[first] = true
	}
	return false
}

func main() {
	fmt.Println(findWhetherExistsPath(3, [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}}, 0, 2))
	fmt.Println(findWhetherExistsPath(5, [][]int{{0, 1}, {0, 2}, {0, 4}, {0, 4}, {0, 1}, {1, 3}, {1, 4}, {1, 3}, {2, 3}, {3, 4}}, 0, 4))
	fmt.Println(findWhetherExistsPath(12, [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 10}, {1, 11}, {1, 4}, {2, 4}, {2, 6}, {2, 9}, {2, 10}, {2, 4}, {2, 5}, {2, 10}, {3, 7}, {3, 7}, {4, 5}, {4, 11}, {4, 11}, {4, 10}, {5, 7}, {5, 10}, {6, 8}, {7, 11}, {8, 10}}, 2, 3))
}
