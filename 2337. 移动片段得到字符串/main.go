package main

import "fmt"

// https://leetcode.cn/problems/move-pieces-to-obtain-a-string/

func canChange(start string, target string) bool {
	i, j := 0, 0
	for i < len(start) && j < len(target) {
		for i < len(start) && start[i] == '_' {
			i++
		}
		for j < len(target) && target[j] == '_' {
			j++
		}
		if i < len(start) && j < len(target) {
			if start[i] != target[j] {
				return false
			}
			if start[i] == 'L' && i < j {
				return false
			}
			if start[i] == 'R' && i > j {
				return false
			}
			j++
			i++
		}
	}

	for i < len(start) && start[i] == '_' {
		i++
	}
	for j < len(target) && target[j] == '_' {
		j++
	}
	return i == len(start) && j == len(target)
}

func main() {
	fmt.Println(canChange("_L__R__R_", "L______RR"))
	fmt.Println(canChange("_L__R", "___RR"))
	fmt.Println(canChange("_R", "R_"))
	fmt.Println(canChange("__", "__"))
}
