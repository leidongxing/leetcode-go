package dp

import "testing"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	} else {
		n := len(nums)
		dp := make([]int, n+1)
		dp[0] = 0
		dp[1] = nums[0]
		for i := 2; i < n+1; i++ {
			v1 := dp[i-1]
			v2 := dp[i-2] + nums[i-1]
			if v1 > v2 {
				dp[i] = v1
			} else {
				dp[i] = v2
			}
		}

		if dp[n] > dp[n-1] {
			return dp[n]
		} else {
			return dp[n-1]
		}
	}
}

func Test_198(t *testing.T) {
	t.Log(rob([]int{2, 1, 1, 2}))    //4
	t.Log(rob([]int{}))              //0
	t.Log(rob([]int{1, 2, 3, 1}))    //4
	t.Log(rob([]int{2, 7, 9, 3, 1})) //12
}
