package dynamic_programming

import "testing"

//乘积最大连续子数组
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//以当前为后缀的连续子数组的长度最大值
	dp := make([]int, len(nums))
	//以当前为后缀的连续子数组的长度最小值
	min := make([]int, len(nums))
	max := nums[0]
	dp[0], min[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		next1 := dp[i-1] * nums[i]
		next2 := min[i-1] * nums[i]
		if next1 > next2 {
			if nums[i] > next1 {
				dp[i] = nums[i]
				min[i] = next2
			} else {
				dp[i] = next1
				if next2 > nums[i] {
					min[i] = nums[i]
				} else {
					min[i] = next2
				}
			}
		} else {
			if nums[i] > next2 {
				dp[i] = nums[i]
				min[i] = next1
			} else {
				dp[i] = next2
				if next1 > nums[i] {
					min[i] = nums[i]
				} else {
					min[i] = next1
				}
			}
		}

		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func Test_152(t *testing.T) {
	t.Log(maxProduct([]int{2, 3, -2, 4})) //6
	t.Log(maxProduct([]int{-2, 0, -1}))   //0
	t.Log(maxProduct([]int{-2, 3, -4}))   //24
}
