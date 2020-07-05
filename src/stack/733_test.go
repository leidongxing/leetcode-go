package stack

import "testing"

//每日温度  单调栈 存的是数组下标  栈顶元素保持最小
func dailyTemperatures(T []int) []int {
	result := make([]int, len(T))
	var stack []int
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			//出栈
			stack = stack[:len(stack)-1]
		}
		//入栈
		stack = append(stack, i)
	}
	return result
}

func Test_733(t *testing.T) {
	t.Log(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
