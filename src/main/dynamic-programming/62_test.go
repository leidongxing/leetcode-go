package dynamic_programming

import "testing"

//不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for j := 1; j < m; j++ {
		for k := 1; k < n; k++ {
			dp[j][k] = dp[j-1][k] + dp[j][k-1]
		}
	}
	return dp[m-1][n-1]
}

func Test_62(t *testing.T) {
	t.Log(uniquePaths(3, 2)) //3
	t.Log(uniquePaths(7, 3)) //28
}
