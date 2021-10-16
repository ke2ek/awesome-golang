package strings

import "strings"

// Palindrome is the word which reads the same backward as forward.
func IsPalindrome(s string) bool {
	s = strings.ToUpper(s)
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
	/*
		m := len(s) / 2
		if len(s)%2 == 0 {
			m--
		}
		for i := 0; i <= m; i++ {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
		return true
	*/
}
