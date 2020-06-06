package dynamic_programming

import "testing"

//买卖股票的最佳时机  买入和卖出一支股票一次
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if buy > prices[i] {
			buy = prices[i]
		} else if prices[i]-buy > profit {
			//利润必须要先买后卖
			profit = prices[i] - buy
		}
	}
	if profit > 0 {
		return profit
	} else {
		return 0
	}
}

func Test_121(t *testing.T) {
	t.Log(maxProfit1([]int{7, 1, 5, 3, 6, 4})) //5
	t.Log(maxProfit1([]int{7, 6, 4, 3, 1}))    //0
}
