package dynamic_programming

import "testing"

//戳气球
func maxCoins(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(nums)
	//从0到i的区间最大值
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	newNums := make([]int, n+2)
	newNums[0], newNums[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		newNums[i] = nums[i-1]
	}
	//i是左侧的
	for i := n - 1; i >= 0; i-- {
		//j是右侧的
		for j := i + 2; j <= n+1; j++ {
			//k是中间的
			for k := i + 1; k < j; k++ {
				sum := newNums[i] * newNums[k] * newNums[j]
				sum += dp[i][k] + dp[k][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

func maxCoins1(nums []int) int {
	//构造新数组
	n := len(nums) + 2
	newNums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		newNums[i+1] = nums[i]
	}
	newNums[0], newNums[n-1] = 1, 1

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		//初始化为-1
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return solve(0, n-1, newNums, dp)
}

//递归的解决
func solve(left, right int, nums []int, dp [][]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if left >= right-1 {
		return 0
	}
	//已经计算出来 不再重复计算
	if dp[left][right] != -1 {
		return dp[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := nums[left] * nums[i] * nums[right]
		sum += solve(left, i, nums, dp) + solve(i, right, nums, dp)
		dp[left][right] = max(dp[left][right], sum)
	}
	return dp[left][right]
}

func Test_312(t *testing.T) {
	t.Log(maxCoins([]int{3, 1, 5, 8})) //167
}
