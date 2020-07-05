package dynamic_programming

import (
	"math"
	"testing"
)

//石子游戏
func stoneGame(piles []int) bool {
	n := len(piles)
	//dp[i][j]表示i到j之间Alex最多可以赢Lee的分数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = piles[i]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			//选用piles[j]与 Lee的选择dp[j+1][i+j]的值
			dp[j][i+j] = int(math.Max(float64(piles[j]-dp[j+1][i+j]), float64(piles[i+j]-dp[j][i+j-1])))
		}
	}
	return dp[0][n-1] > 0
}

func Test_377(t *testing.T) {
	t.Log(stoneGame([]int{5, 3, 4, 5}))  //true
	t.Log(stoneGame([]int{3, 2, 10, 4})) //true
}
