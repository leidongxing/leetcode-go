package dynamic_programming

import "testing"

//暴力破解
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := INT_MIN
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		if max < sum {
			max = sum
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

//连续子数组的最大和
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < n; i++ {
		//如果之前是负数 就不要了
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		//更新最大值
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func Test_53(t *testing.T) {
	//var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} //6
	//t.Log(maxSubArray(nums))
	t.Log(maxSubArray([]int{-2, 1})) //1
}
