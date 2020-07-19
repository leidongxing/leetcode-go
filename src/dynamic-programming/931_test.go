package dynamic_programming

import (
	"math"
	"testing"
)

//下降路径最小和
func minFallingPathSum(A [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make([]int, len(A[0]))
	}
	for i := 0; i < len(A[0]); i++ {
		dp[0][i] = A[0][i]
	}

	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			dp[i][j] = A[i][j]
			last := math.MaxInt32
			for k := 0; k < len(A[0]); k++ {
				if j-k > 1 || k-j > 1 {
					continue
				}
				last = min(last, dp[i-1][k])
			}
			dp[i][j] += last
		}
	}

	result := math.MaxInt32
	for i := 0; i < len(A[0]); i++ {
		result = min(result, dp[len(A)-1][i])
	}
	return result
}
func Test_931(t *testing.T) {
	t.Log(minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             //12
	t.Log(minFallingPathSum([][]int{{-80, -13, 22}, {83, 94, -5}, {73, -48, 61}})) //-66
}
