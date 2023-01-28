package main

import "fmt"

// https://leetcode.cn/problems/strong-password-checker-ii/
// 这个代码也没有什么好解释的，非常简答的一道题
func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	var little, big, num, special bool
	var disSame = true
	for i, c := range password {
		if 'a' <= c && c <= 'z' {
			little = true
		} else if 'A' <= c && c <= 'Z' {
			big = true
		} else if '0' <= c && c <= '9' {
			num = true
		} else {
			special = true
		}
		if i > 0 && password[i] == password[i-1] {
			disSame = false
		}
	}
	return little && big && num && special && disSame
}

func main() {
	for _, s := range []string{
		"abcde",
	} {
		fmt.Println(s, strongPasswordCheckerII(s))
	}
}
