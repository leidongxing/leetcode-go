package dynamic_programming

import (
	"testing"
)

//零钱兑换
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return helper(coins, amount, make([]int, amount))
}

//dp[] 拼凑 钱数对应的
func helper(coins []int, target int, dp []int) int {
	if target < 0 {
		return -1
	} else if target == 0 {
		return 0
	} else {
		if dp[target-1] != 0 {
			return dp[target-1]
		}
		IntMax := int(^uint(0) >> 1)
		min := IntMax
		for _, coin := range coins {
			res := helper(coins, target-coin, dp)
			if res >= 0 && res < min {
				min = 1 + res
			}
		}
		if min == IntMax {
			dp[target-1] = -1
		} else {
			dp[target-1] = min
		}
		return dp[target-1]
	}
}

func Test_322(t *testing.T) {
	t.Log(coinChange([]int{1, 2, 5}, 11)) //3
	t.Log(coinChange([]int{2}, 3))        //-1
}
