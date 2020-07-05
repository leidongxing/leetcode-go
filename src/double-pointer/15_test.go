package double_pointer

import (
	"sort"
	"testing"
)

//三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	//排序
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		//排除重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < n; j++ {
			//排除重复元素
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			if j == k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return result
}

func Test_15(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
