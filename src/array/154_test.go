package array

import "testing"

//寻找旋转排序数组中的最小值 II
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}
	return nums[0]

	//left, right := 0, len(nums)-1
	//for left < right {
	//	mid := (left + right) / 2
	//	if nums[mid] > nums[left] {
	//		if nums[mid] > nums[right] {
	//			left = mid
	//		} else if nums[mid] < nums[right] {
	//			right = mid
	//		} else {
	//			return nums[mid]
	//		}
	//	} else if nums[mid] < nums[left] {
	//		if nums[mid] < nums[right] {
	//			right = mid
	//		} else if nums[mid] > nums[right] {
	//			left = mid
	//		} else {
	//			return nums[mid]
	//		}
	//	}
	//}
	//return -1
}

func Test_154(t *testing.T) {
	t.Log(findMin([]int{3, 4, 5, 1, 2}))
	t.Log(findMin([]int{2, 2, 2, 0, 1}))
}
