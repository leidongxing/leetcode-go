package array

import (
	"testing"
)

//计算右侧小于当前元素的个数
//暴力方法
func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, len(nums))
	result[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func Test_315(t *testing.T) {
	//t.Log(countSmaller([]int{5,2,6,1}))//[2 1 1 0]
	t.Log(countSmaller([]int{})) //[]
}
