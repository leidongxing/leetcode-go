package interview

import "testing"

//恢复空格
func respace(dictionary []string, sentence string) int {
	mySet := map[string]bool{}
	for _, d := range dictionary {
		mySet[d] = true
	}
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(sentence)
	//前i个字符中的最少匹配数
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		//初始化 i-1不匹配 第i个也不会匹配
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if mySet[sentence[j:i]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}

func Test_17_13(t *testing.T) {
	t.Log(respace([]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother")) //7
}
