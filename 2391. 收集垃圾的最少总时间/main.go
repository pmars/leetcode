package main

import "fmt"

// https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage/

/*
分析：
	这个题理解起来一点不麻烦，也就是计算上麻烦一点，其他的很简单
*/
func garbageCollection(garbage []string, travel []int) int {
	var maxIndexM, maxIndexG, maxIndexP int // 定义三个变量，记录最大的位置，处理垃圾车运行的距离
	var sum, sumTravel int                  // 定义每种垃圾的回收时间

	for i, g := range garbage {
		for j := range g {
			switch g[j] {
			case 'M':
				maxIndexM = i
			case 'G':
				maxIndexG = i
			case 'P':
				maxIndexP = i
			}
		}
		sum += len(g)
	}
	// fmt.Println(sum, maxIndexM, maxIndexG, maxIndexP)

	// 这边遍历所有的距离，是因为最后一个地方肯定有垃圾，垃圾车肯定需要计算出这个结果的
	for i := 0; i < len(travel); i++ {
		sumTravel += travel[i]
		if i+1 == maxIndexP {
			sum += sumTravel
		}
		if i+1 == maxIndexG {
			sum += sumTravel
		}
		if i+1 == maxIndexM {
			sum += sumTravel
		}
	}

	return sum
}

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))
}
