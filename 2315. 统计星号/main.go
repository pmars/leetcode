package main

import "fmt"

// https://leetcode.cn/problems/count-asterisks/

func countAsterisks(s string) int {
	times := 0
	check := true

	for _, c := range s {
		if c == '|' {
			check = !check
		} else if c == '*' && check {
			times++
		}
	}
	return times
}

func main() {
	for _, s := range []string{
		"l|*e*et|c**o|*de|",
		"iamprogrammer",
		"yo|uar|e**|b|e***au|tifu|l",
	} {
		fmt.Println(s, countAsterisks(s))
	}
}
