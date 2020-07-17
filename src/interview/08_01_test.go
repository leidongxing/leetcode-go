package interview

import "testing"

//三步问题
func waysToStep(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	} else {
		dp := make([]int, n)
		dp[0] = 1
		dp[1] = 2
		dp[2] = 4
		for i := 3; i < n; i++ {
			dp[i] = (dp[i-3] + dp[i-2] + dp[i-1]) % 1000000007
		}
		return dp[n-1]
	}
}

func Test_08_01(t *testing.T) {
	t.Log(waysToStep(1))
	t.Log(waysToStep(2))
	t.Log(waysToStep(61))
	t.Log(waysToStep(76))
}
