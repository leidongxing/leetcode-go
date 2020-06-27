package array

import "testing"

//缺失的第一个正数
func firstMissingPositive(nums []int) int {
	for i := 1; i <= len(nums); i++ {
		isFound := false
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				isFound = true
				break
			}
		}
		if !isFound {
			return i
		}
	}
	return len(nums) + 1
}

func Test_41(t *testing.T) {
	t.Log(firstMissingPositive([]int{1, 2, 3}))         //4
	t.Log(firstMissingPositive([]int{1, 2, 0}))         //3
	t.Log(firstMissingPositive([]int{3, 4, -1, 1}))     //2
	t.Log(firstMissingPositive([]int{7, 8, 9, 11, 12})) //1

}
