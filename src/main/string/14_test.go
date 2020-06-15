package string

import "testing"

//最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	} else {
		minLength := len(strs[0])
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) == 0 {
				return ""
			} else if len(strs[i]) < minLength {
				minLength = len(strs[i])
			}
		}
		j := 0
	I:
		for ; j < minLength; j++ {
			for k := 1; k < len(strs); k++ {
				if strs[k][j] != strs[0][j] {
					break I
				}
			}
		}
		return strs[0][0:j]
	}
}

func Test_14(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"})) //fl
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))    //
}
