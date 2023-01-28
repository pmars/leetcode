package main

import "fmt"

/*
分析：
	这个题的限制是只删掉一个元素，这也方便了计算了
	直接一个for，之后算一下前面的内容，和后面的内容的和，对比一下是否相同就可以了
	但，在想一下，for的过程中，计算前后的内容和，是不是可以提前算出来，sum(奇)，sum(偶)，之后在for的过程中，动态的改变这个值的内容，
	并记录前面的内容和，这样在for的过程中，就只需要加减法计算了
*/

func waysToMakeFair(nums []int) int {
	sumJiQian, sumJiHou, sumOuQian, sumOuHou := 0, 0, 0, 0
	times := 0
	for i, n := range nums {
		if i%2 == 0 {
			sumOuHou += n
		} else {
			sumJiHou += n
		}
	}

	for i, n := range nums {
		// 将第 i 个数字删除
		if i%2 == 0 {
			sumOuHou -= n
		} else {
			sumJiHou -= n
		}

		// 删除之后，后面的奇数index变为了偶数，偶数变为奇数
		// 求和的时候，交叉相加即可
		if sumJiQian+sumOuHou == sumOuQian+sumJiHou {
			times++
		}

		// 给刚才删除的数字，加回来，变为了前面的数字，偶数和奇数根据 i 来限定即可
		if i%2 == 0 {
			sumOuQian += n
		} else {
			sumJiQian += n
		}
	}
	return times
}

func main() {
	for _, s := range [][]int{
		{1, 2, 3},
		{2, 1, 6, 4},
		{6, 1, 7, 4, 1},
	} {
		fmt.Println(s, waysToMakeFair(s))
	}
}
