package string

import (
	"strings"
)

func longestPalindrome(s string) string {
	if isSame(s) {
		return s
	}
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindromic(s) {

			}

		}
	}
	return result.String()
}

func isSame(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[0] != s[i] {
			return false
		}
	}
	return true
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
