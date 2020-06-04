package dynamic_programming

import "testing"

//最长上升子序列 动态规划
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		cur := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] > cur {
					cur = dp[j]
				}
			}
		}
		//为之前所有dp中的最大值+1
		dp[i] = cur + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func Test_300(t *testing.T) {
	t.Log(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) //4
	t.Log(lengthOfLIS([]int{4, 10, 3, 8, 9}))             //3
	t.Log(lengthOfLIS([]int{2, 2}))                       //1
}
