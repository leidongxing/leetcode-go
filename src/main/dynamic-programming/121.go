package dynamic_programming

func maxProfit(prices []int) int {
	var max int = 0
	var paid int = 1008611
	for i := 0; i < len(prices); i++ {
		if paid > prices[i] {
			paid = prices[i]
		} else if max < prices[i]-paid {
			max = prices[i] - paid
		}
	}
	return max
}
