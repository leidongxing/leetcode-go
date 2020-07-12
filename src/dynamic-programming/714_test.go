package dynamic_programming

import "testing"

//买卖股票的最佳时机含手续费
func maxProfit3(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[len(prices)-1][1]
}

func Test_714(t *testing.T) {
	t.Log(maxProfit3([]int{1, 3, 2, 8, 4, 9}, 2)) //8
}
