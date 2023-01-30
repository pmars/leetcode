package main

import "fmt"

// https://leetcode.cn/problems/string-without-aaa-or-bbb/

func strWithout3a3b(a int, b int) string {
	list := make([]uint8, a+b)
	aTimes, bTimes := 0, 0
	for i := 0; a > 0 || b > 0; i++ {
		aFlag := true
		if a > b {
			if aTimes >= 2 {
				aFlag = false
			}
		} else {
			if bTimes < 2 {
				aFlag = false
			}
		}
		if aFlag {
			aTimes++
			bTimes = 0
			list[i] = 'a'
			a--
		} else {
			bTimes++
			aTimes = 0
			list[i] = 'b'
			b--
		}
	}
	return string(list)
}
func main() {
	fmt.Println(strWithout3a3b(3, 4))
	fmt.Println(strWithout3a3b(1, 1))
	fmt.Println(strWithout3a3b(4, 1))
	fmt.Println(strWithout3a3b(1, 4))
}
