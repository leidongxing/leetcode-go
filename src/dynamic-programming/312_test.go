package dynamic_programming

import "testing"

//戳气球
func maxCoins(nums []int) int {
	n := len(nums) + 2
	newNums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		newNums[i+1] = nums[i]
	}
	//边界处理
	newNums[0], newNums[n-1] = 1, 1
	memo := make([][]int, n, n)
	return dp(memo, 1, 1)
}

func dp(memo [][]int, left int, right int) int {
	if left+1 == right {
		return 0
	}
	if memo[left][right] > 0 {
		return memo[left][right]
	}
	ans := 0
	for i := left + 1; i < right; i++ {

	}
	return ans
}

func Test_312(t *testing.T) {
	t.Log(maxCoins([]int{3, 1, 5, 8})) //167
}
