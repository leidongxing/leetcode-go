package dynamic_programming

import (
	"math"
	"testing"
)

//下降路径最小和II
func minFallingPathSum2(arr [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, len(arr[0]))
	}
	for i := 0; i < len(arr[0]); i++ {
		dp[0][i] = arr[0][i]
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			dp[i][j] = arr[i][j]
			last := math.MaxInt32
			for k := 0; k < len(arr[0]); k++ {
				if j == k {
					continue
				}
				last = min(last, dp[i-1][k])
			}
			dp[i][j] += last
		}
	}

	result := math.MaxInt32
	for i := 0; i < len(arr[0]); i++ {
		result = min(result, dp[len(arr)-1][i])
	}
	return result
}

func Test_1289(t *testing.T) {
	t.Log(minFallingPathSum2([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) //13
}
