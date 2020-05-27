package dynamic_programming

import "testing"

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func Test_70(t *testing.T) {
	t.Log(climbStairs(2))
	t.Log(climbStairs(3))
}
