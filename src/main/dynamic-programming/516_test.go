package dynamic_programming

import "testing"

//最长回文子序列
func longestPalindromeSubseq(s string) int {
	//二维数组保存 从i到j 最长的回文子序列长度
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	//向左扩张
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		//向右扩张
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	//最后返回 从0开头到结尾的
	return dp[0][len(s)-1]
}

func Test_516(T *testing.T) {
	T.Log(longestPalindromeSubseq("bbbab")) //4
	T.Log(longestPalindromeSubseq("cbbd"))  //2
}
