package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	indexMap map[int][]int
	r        *rand.Rand
}

func Constructor(nums []int) Solution {
	indexMap := make(map[int][]int) // index list map
	for i, n := range nums {
		if _, ok := indexMap[n]; !ok {
			indexMap[n] = make([]int, 0)
		}
		indexMap[n] = append(indexMap[n], i)
	}

	return Solution{
		indexMap: indexMap,
		r:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Pick(target int) int {
	list, ok := this.indexMap[target]
	if !ok {
		return 0
	}
	if len(list) == 1 {
		return list[0]
	}

	return list[this.r.Intn(len(list))]
}

func main() {
	obj := Constructor([]int{0, 1, 2, 3, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6})
	for i := 0; i < 100; i++ {
		fmt.Println(obj.Pick(5))
	}
}
