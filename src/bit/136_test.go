package bit

import "testing"

//异或运算  111^111=0  111^0=111
func singleNumber(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

func Test_136(t *testing.T) {
	t.Log(singleNumber([]int{2, 2, 1}))
	t.Log(singleNumber([]int{4, 1, 2, 1, 2}))
}
