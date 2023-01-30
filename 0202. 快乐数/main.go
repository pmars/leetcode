package main

import "fmt"

// https://leetcode.cn/problems/happy-number/

// 设置一个map记录存在过的内容，如果遇到了，那么就是循环了
func isHappyV1(n int) bool {
	existMap := map[int]bool{
		n: true,
	}

	for {
		n = getNext(n)
		if n == 1 {
			return true
		}
		if existMap[n] {
			return false
		}
		existMap[n] = true
	}
}

// 看评论区，有人说快慢指针的方法，试一下
// 过了，这个果然凶残，速度极快
func isHappy(n int) bool {
	slow := getNext(n)
	fast := getNext(slow)
	for fast != slow {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		m := n % 10
		sum += m * m
		n /= 10
	}
	return sum
}

func main() {
	for _, i := range []int{19, 25, 2} {
		fmt.Println(i, isHappy(i))
	}
}
