package string

import "testing"

//删除回文子序列
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	} else if isPalindrome(s) {
		return 1
	} else {
		return 2
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func Test_1332(t *testing.T) {
	t.Log(isPalindrome("ababa"))
	t.Log(isPalindrome("abba"))
}
