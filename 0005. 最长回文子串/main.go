package main

import "fmt"

// https://leetcode.cn/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	start, end := 0, -1
	for i := range s {
		// 以 i 为中心
		j := 1
		for i-j >= 0 && i+j < len(s) && s[i-j] == s[i+j] {
			j++
		}
		startT, endT := i-j+1, i+j-1
		if endT-startT > end-start {
			end, start = endT, startT
		}

		// 以 i,i+1 为中心
		if i+1 < len(s) && s[i] == s[i+1] {
			j = 1
			for i-j >= 0 && i+1+j < len(s) && s[i-j] == s[i+1+j] {
				j++
			}
			startT, endT := i-j+1, i+1+j-1
			if endT-startT > end-start {
				end, start = endT, startT
			}
		}
	}
	return s[start : end+1]
}

func main() {
	for _, s := range []string{
		"abcbcb",
		"abcb",
		"abccb",
	} {
		fmt.Println(s, longestPalindrome(s))
	}
}
