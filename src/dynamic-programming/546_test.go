package dynamic_programming

import (
	"testing"
)

//移除盒子
func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	return calculatePoints(boxes, dp, 0, len(boxes)-1, 0)
}

func calculatePoints(boxes []int, dp [100][100][100]int, l, r, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	if l > r {
		return 0
	}
	if dp[l][r][k] != 0 {
		return dp[l][r][k]
	}
	for r > l && boxes[r] == boxes[r-1] {
		r--
		k++
	}
	dp[l][r][k] = calculatePoints(boxes, dp, l, r-1, 0) + (k+1)*(k+1)
	for i := l; i < r; i++ {
		if boxes[i] == boxes[r] {
			dp[l][r][k] = max(dp[l][r][k], calculatePoints(boxes, dp, l, i, k+1)+calculatePoints(boxes, dp, i+1, r-1, 0))
		}
	}
	return dp[l][r][k]
}

func Test_546(t *testing.T) {
	t.Log(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}
