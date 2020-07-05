package stack

import "testing"

//对于每一个高度 求出它的左右边界 即为面积
//如果有两根柱子 j0 j1  如果j1<j0 j0会被j1挡住
//单调栈  柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	monoStack := []int{}
	//从左往右  找出每一个位置的左边界 index
	for i := 0; i < n; i++ {
		//如果元素高度始终小于栈顶元素  入栈
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		//为空时 设置哨兵
		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	//从右向左 找到每一个位置的右边界 index
	for i := n - 1; i >= 0; i-- {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			right[i] = n
		} else {
			right[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Test_84(t *testing.T) {
	t.Log(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) //10
	t.Log(largestRectangleArea([]int{2, 0, 2}))          //2
	t.Log(largestRectangleArea([]int{2, 1, 2}))          //3

}
