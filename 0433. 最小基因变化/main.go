package main

import "fmt"

// https://leetcode.cn/problems/minimum-genetic-mutation/

/*
分析：
	乍一看，两个字符串变化，直接计算两个字符串不通的字符就可以了，但不然
	字符串变化的过程中的状态，是需要在bank里面的，而贪心的从一端到另一端的变化，不一定在bank内，所以这块还得做一个处理
	而且，也不一定一样的字符就不需要变了，可能这个先变化了，之后其他的在变，之后这个在变回来，这么想，还是有点复杂的
	那么这边直接做一个广搜就可以了
*/
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}

	bankMap := make(map[string]bool)        // 这个里面的可以变
	stepList := []string{startGene}         // 基因序列转化过程中的状态，
	doneMap := map[string]int{startGene: 0} // 这个里面的已经变化过了，不用在往回变了

	// 填充bankMap，加速处理
	for _, b := range bank {
		bankMap[b] = true
	}

	if !bankMap[endGene] {
		return -1
	}

	for len(stepList) > 0 {
		start := stepList[0]
		stepList = stepList[1:]

		oneStepGeneList := diffOneStep(bankMap, start)
		for _, g := range oneStepGeneList {
			if _, ok := doneMap[g]; ok {
				continue
			}
			if g == endGene {
				return doneMap[start] + 1
			}
			stepList = append(stepList, g)
			doneMap[g] = doneMap[start] + 1
			// fmt.Println(g, doneMap[g])
		}
	}
	return -1
}

// 返回基因gene变化一步的基因序列，这个序列是在bankMap里面的
// bank长度也就10个，直接计算一个字符不一样的，也算是剪枝了
func diffOneStep(bankMap map[string]bool, gene string) []string {
	resList := make([]string, 0)
	for bankGene := range bankMap {
		diff := 0
		for i := 0; i < len(bankGene); i++ {
			if bankGene[i] != gene[i] {
				diff++
				if diff > 1 {
					break
				}
			}
		}
		if diff == 1 {
			resList = append(resList, bankGene)
		}
	}
	return resList
}

func main() {
	fmt.Println(minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))                                                             // 1
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))                                     // 2
	fmt.Println(minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))                                     // 3
	fmt.Println(minMutation("AAAAAAAA", "AAAAACGG", []string{"AAAAAAGA", "AAAAAGGA", "AAAAACGA", "AAAAACGG", "AAAAAAGG", "AAAAAAGC"})) // 3
}
