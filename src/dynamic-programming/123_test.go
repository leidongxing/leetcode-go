package dynamic_programming

import (
	"testing"
)

//买卖股票的最佳时机 III  最多可以完成两笔交易
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		// 0 尚未买入  1 买入一次  2卖出一次  3买入两次  4卖出两次
		dp[i] = make([]int, 5)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = 0
	dp[0][4] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = 0
		if dp[i-1][1] > dp[i-1][0]-prices[i] {
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][1] = dp[i-1][0] - prices[i]
		}

		if dp[i-1][2] > dp[i-1][1]+prices[i] {
			dp[i][2] = dp[i-1][2]
		} else {
			dp[i][2] = dp[i-1][1] + prices[i]
		}

		if dp[i-1][3] > dp[i-1][2]-prices[i] {
			dp[i][3] = dp[i-1][3]
		} else {
			dp[i][3] = dp[i-1][2] - prices[i]
		}

		if dp[i-1][4] > dp[i-1][3]+prices[i] {
			dp[i][4] = dp[i-1][4]
		} else {
			dp[i][4] = dp[i-1][3] + prices[i]
		}
	}
	max := func(a int, b int, c int) int {
		if a > b && a > c {
			return a
		} else if b > a && b > c {
			return b
		} else {
			return c
		}
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][2], dp[len(prices)-1][4])
}

func Test_123(t *testing.T) {
	//t.Log(maxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4})) //6
	//t.Log(maxProfit3([]int{1, 2, 3, 4, 5})) //4
	t.Log(maxProfit3([]int{7, 6, 4, 3, 1})) //0
}
