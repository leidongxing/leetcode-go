package dynamic_programming

//最佳买卖股票时机含冷冻期
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		//前一天不在冷冻期今天买入
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		//持有股票卖掉
		f[i][1] = f[i-1][0] + prices[i]

		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	//最后肯定是 不持有股票
	return max(f[n-1][1], f[n-1][2])
}
