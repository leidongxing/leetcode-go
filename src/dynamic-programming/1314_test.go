package dynamic_programming

import "testing"

func matrixBlockSum(mat [][]int, K int) [][]int {
	dp := make([][]int, len(mat))
	for i := 0; i < len(mat[0]); i++ {
		dp[i] = make([]int, len(mat[0]))
	}
	dp[0][0] = mat[0][0] + base(mat, K, 0, 0)

	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			dp[i][j] = dp[i-1][j-1] + base(mat, K, i, j)
		}
	}
	return dp
}

func base(mat [][]int, K int, m int, n int) int {
	sum := 0
	for i := m; i <= K; i++ {
		sum += mat[i][m]
		for j := n; j <= K; j++ {
			sum += mat[n][j]
		}
	}
	return sum
}

func Test_1314(t *testing.T) {
	t.Log(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1)) //
}
