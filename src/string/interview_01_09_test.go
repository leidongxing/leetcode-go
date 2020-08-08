package string

import "testing"

//字符串轮转
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else if len(s1) == 0 {
		return true
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[0] {
			leftResult, j := flipLeft(s1, s2, i)
			if !leftResult {
				continue
			}
			rightResult := flipRight(s1, s2, j)
			if rightResult {
				return rightResult
			}
		}
	}
	return false
}

func flipLeft(s1 string, s2 string, i int) (bool, int) {
	j := 0
	for i+j < len(s1) {
		if s1[i+j] != s2[j] {
			return false, -1
		}
		j++
	}
	return true, j
}

func flipRight(s1 string, s2 string, j int) bool {
	k := 0
	for j < len(s2) {
		if s1[k] != s2[j] {
			return false
		} else {
			k++
			j++
		}
	}
	return true
}

func Test_01_09(t *testing.T) {
	t.Log(isFlipedString("waterbottle", "erbottlewat"))
	t.Log(isFlipedString("aa", "aba"))
	t.Log(isFlipedString("", ""))
	t.Log(isFlipedString("aba", "bab"))
}
