package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		if len(p) == 0 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < len(p)-1; i++ {
		if p[i] == '.' && p[i+1] == '*' {
			return true
		}
	}

	flag := true
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if len(p) <= i+j {
				return false
			}
			if p[i+j] != s[j] {
				if p[i+j] == '.' {
					continue
				}
				if p[i+j] == '*' {

				}
			}
		}
	}
	return flag
}
