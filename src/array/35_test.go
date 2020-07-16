package array

import "testing"

//搜索插入位置
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func Test_35(t *testing.T) {
	t.Log(searchInsert([]int{1, 3, 5, 6}, 5)) //2
	t.Log(searchInsert([]int{1, 3, 5, 6}, 2)) //1
	t.Log(searchInsert([]int{1, 3, 5, 6}, 7)) //4
	t.Log(searchInsert([]int{1, 3, 5, 6}, 0)) //0
	t.Log(searchInsert([]int{1}, 1))          //0
}
