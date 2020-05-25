package dp

import "testing"

//使用最小花费爬楼梯
//dp[i]=Min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-1])
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1, n+1)
	dp[0] = 0
	dp[1] = cost[0]
	dp[2] = cost[1]
	for i := 3; i < n+1; i++ {
		if dp[i-1]+cost[i-1] < dp[i-2]+cost[i-1] {
			dp[i] = dp[i-1] + cost[i-1]
		} else {
			dp[i] = dp[i-2] + cost[i-1]
		}
	}
	if dp[n-1] < dp[n] {
		return dp[n-1]
	} else {
		return dp[n]
	}
}

func Test_746(t *testing.T) {
	t.Log(minCostClimbingStairs([]int{10, 15, 20}))                         //15
	t.Log(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) //6
	t.Log(minCostClimbingStairs([]int{0, 1, 2, 0}))                         //1
}
