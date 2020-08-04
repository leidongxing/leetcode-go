package string

import "testing"

//回文排列
func canPermutePalindrome(s string) bool {
	hashmap := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
	}
	count := 0
	for _, w := range hashmap {
		if w%2 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return count <= 1
}

func Test_01_04(t *testing.T) {
	t.Log(canPermutePalindrome("tactcoa"))

}
