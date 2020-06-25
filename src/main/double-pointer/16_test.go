package double_pointer

import (
	"math"
	"sort"
	"testing"
)

//最接近的三数之和1
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}

	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	for i := 0; i < n; i++ {
		//排除重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			update(sum)
			if sum > target {
				k0 := k - 1
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func Test_16(t *testing.T) {
	t.Log(threeSumClosest([]int{-1, 2, 1, -4}, 1)) //2
	t.Log(threeSumClosest([]int{0, 0, 0}, 1))      //0
}
