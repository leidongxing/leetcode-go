package array

import "testing"

//除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	answer[0] = 1
	//计算当前index左侧乘积
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		//累乘右侧乘积
		R *= nums[i]
	}
	return answer
}

func Test_238(t *testing.T) {
	t.Log(productExceptSelf([]int{1, 2, 3, 4}))
}
