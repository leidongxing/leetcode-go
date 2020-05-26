package dynamic_programming

import "testing"

func minPathSum(grid [][]int) int {
	//初始化数组
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for k := 1; k < m; k++ {
		dp[k][0] = dp[k-1][0] + grid[k][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func Test_64(t *testing.T) {
	//t.Log(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	t.Log(minPathSum([][]int{{1, 2, 5}, {3, 2, 1}}))
}
