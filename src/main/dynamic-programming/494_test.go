package dynamic_programming

import (
	"testing"
)

//目标和  sum(P)-sum(N)=target    sum(P)+sum(N)+sum(P)-sum(N)=sum(P)+sum(N)+target  2*sum(P)=target+sum(N)
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	//和小于S   sum+S其实就是2*sum(P)
	if sum < S || (sum+S)&1 == 1 {
		return 0
	}
	target := (sum + S) / 2
	//容量为target的有多少种方案
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[target]
}

//目标和  从前往后递归
var result = 0

func findTargetSumWays2(nums []int, S int) int {
	result = 0
	calculate(nums, 0, 0, S)
	return result
}

func calculate(nums []int, index int, sum int, S int) {
	if index == len(nums) {
		if sum == S {
			result++
		}
	} else {
		calculate(nums, index+1, sum+nums[index], S)
		calculate(nums, index+1, sum-nums[index], S)
	}
}

//目标和  从后往前递归
func findTargetSumWays1(nums []int, S int) int {
	return findTarget(nums, S, len(nums)-1)
}

func findTarget(nums []int, target int, index int) int {
	if index == 0 {
		if nums[index] == 0 && target == 0 {
			return 2
		} else if nums[index] == target || nums[index] == -target {
			return 1
		} else {
			return 0
		}
	}
	nextIndex := index - 1
	return findTarget(nums, target-nums[index], nextIndex) + findTarget(nums, target+nums[index], nextIndex)
}

func Test_494(t *testing.T) {
	t.Log(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))             //5
	t.Log(findTargetSumWays([]int{1, 0}, 1))                      //2
	t.Log(findTargetSumWays([]int{0, 1}, 1))                      //2
	t.Log(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1)) //256
}
