package main

import "fmt"

// https://leetcode.cn/problems/circular-sentence/

func isCircularSentence(sentence string) bool {
	var i int
	var s uint8
	for i < len(sentence) {
		for i < len(sentence) && sentence[i] != ' ' {
			s = sentence[i]
			i++
		}
		for i < len(sentence) && sentence[i] == ' ' {
			i++
		}
		if i < len(sentence) && sentence[i] != s {
			return false
		}
	}
	return sentence[0] == sentence[len(sentence)-1]
}

func main() {
	for _, s := range []string{
		"abcbcb ba",
		"abcb bc",
		"abccb baaaa",
		"abccbbaaaa",
	} {
		fmt.Println(s, isCircularSentence(s))
	}
}
