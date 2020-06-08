package dynamic_programming

import "testing"

// 买卖股票的最佳时机 II  多次买卖一支股票
// 0持有现金 1持有股票 状态转移 0-1-0-1-0-1-0-1-0
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	//两种状态 用二维数组保留状态
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		//判断 昨天股票今天卖出  昨天现金 今天保持
		if dp[i-1][1]+prices[i] > dp[i-1][0] {
			dp[i][0] = dp[i-1][1] + prices[i]
		} else {
			dp[i][0] = dp[i-1][0]
		}
		//判断 昨天现金今天买入  昨天股票 今天保持
		if dp[i-1][0]-prices[i] > dp[i-1][1] {
			dp[i][1] = dp[i-1][0] - prices[i]
		} else {
			dp[i][1] = dp[i-1][1]
		}
	}
	//返回最后一天持有现金
	return dp[len(prices)-1][0]
}

func Test_122(t *testing.T) {
	t.Log(maxProfit2([]int{7, 1, 5, 3, 6, 4})) //7
	t.Log(maxProfit2([]int{1, 2, 3, 4, 5}))    //4
	t.Log(maxProfit2([]int{7, 6, 4, 3, 1}))    //0
}
