package double_pointer

import (
	"math"
)

//长度最小的子数组  假想一个队列 不断入队 直到和大于s  更新长度
func minSubArrayLen(s int, nums []int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
