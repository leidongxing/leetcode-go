package string

import "testing"

//重复的子字符串
func repeatedSubstringPattern(s string) bool {
	//至少得重复一次
	for i := 0; i < len(s)/2; i++ {
		if len(s)%(i+1) == 0 && check(s, i) {
			return true
		}
	}
	return false
}

func check(s string, end int) bool {
	j := 0
	for i := end + j + 1; i < len(s); i++ {
		if s[i] != s[j] {
			return false
		}
		j++
	}
	return true
}

func Test_439(t *testing.T) {
	t.Log(repeatedSubstringPattern("abab"))         //true
	t.Log(repeatedSubstringPattern("aba"))          //false
	t.Log(repeatedSubstringPattern("abcabcabcabc")) //true
	t.Log(repeatedSubstringPattern("aabaaba"))      //false
}
