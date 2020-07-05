package dynamic_programming

import "testing"

//取 1- n-1  2-n两个dp数组最大值
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		nums1 := nums[0 : len(nums)-1]
		nums2 := nums[1:len(nums)]
		v1 := myRob(nums1)
		v2 := myRob(nums2)
		if v1 > v2 {
			return v1
		} else {
			return v2
		}
	}
}

func myRob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if nums[1] > nums[0] {
		dp[1] = nums[1]
	} else {
		dp[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		v1 := dp[i-1]
		v2 := dp[i-2] + nums[i]
		if v1 > v2 {
			dp[i] = v1
		} else {
			dp[i] = v2
		}
	}
	return dp[len(nums)-1]
}

func Test_213(t *testing.T) {
	t.Log(rob2([]int{0, 0}))
	//t.Log(rob2([]int{2, 1, 1, 2}))    //3
	//t.Log(rob2([]int{}))              //0
	//t.Log(rob2([]int{1, 2, 3, 1}))    //4
	//t.Log(rob2([]int{2, 7, 9, 3, 1})) //11
}
