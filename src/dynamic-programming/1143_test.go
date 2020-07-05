package dynamic_programming

import "testing"

//最长公共子序列
//公共子序列必然在两个串里都有
func longestCommonSubsequence(text1 string, text2 string) int {
	//两个字符串 构造二维dp数组
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(text1)+1; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	max := 0

	dp[0][0] = 0
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}

			if dp[i+1][j+1] > max {
				max = dp[i+1][j+1]
			}
		}
	}

	return max
}

func Test_1143(t *testing.T) {
	t.Log(longestCommonSubsequence("abcde", "ace")) //3
	t.Log(longestCommonSubsequence("abc", "abc"))   //3
	t.Log(longestCommonSubsequence("abc", "def"))   //0
}
