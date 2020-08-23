package string

import "testing"

//回文子串
func countSubstrings(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		//回文串为奇数
		num += count(s, i, i)
		//回文串为偶数
		num += count(s, i, i+1)
	}
	return num
}

func count(s string, start, end int) int {
	num := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		num++
		start--
		end++
	}
	return num
}

func Test_647(t *testing.T) {
	t.Log(countSubstrings("abc"))
	t.Log(countSubstrings("aaa"))
}
