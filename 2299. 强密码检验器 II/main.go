package main

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
