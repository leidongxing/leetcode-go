package array

import (
	"sort"
	"testing"
)

//最长连续序列
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cur++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			cur = 1
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cur++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			cur = 1
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

func Test_128(T *testing.T) {
	T.Log(longestConsecutive([]int{}))                                  //0
	T.Log(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))              //4
	T.Log(longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})) //7
	T.Log(longestConsecutive([]int{1, 2, 0, 1}))                        //3
}
