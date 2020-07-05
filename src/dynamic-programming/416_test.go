package dynamic_programming

import "testing"

//分割等和子集
func canPartition(nums []int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	//一半是奇数 不满足要求
	if sum&1 == 1 {
		return false
	}
	target := sum / 2

	//dp[i][j] 从0到i中选择数 和等于j
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	if nums[0] == target {
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] <= j {
				//i-1个元素中能构成j  或者 i-1个元素能构成j-nums[i] +nums[i]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}

			if dp[i][target] {
				return true
			}
		}
	}
	return dp[len(nums)-1][target]
}

func Test_416(t *testing.T) {
	t.Log(canPartition([]int{1, 5, 11, 5})) //true
	t.Log(canPartition([]int{1, 2, 3, 5}))  //false
}
