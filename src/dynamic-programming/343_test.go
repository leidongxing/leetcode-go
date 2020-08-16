package dynamic_programming

import "testing"

//剪绳子
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 2; i <= n; i++ {
		currentMax := 0
		for j := 1; j < i; j++ {
			//i-j 继续拆分
			currentMax = max(currentMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = currentMax
	}
	return dp[n]
}

func Test_343(t *testing.T) {
	t.Log(cuttingRope(10))
}
