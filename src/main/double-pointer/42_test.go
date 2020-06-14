package double_pointer

import "testing"

//接雨水  water[i]=min( max(height[0...i],  max(height[i...end])) -height[i]
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	sum := 0
	leftMax := height[0]
	rightMax := height[len(height)-1]
	left := 0
	right := len(height) - 1

	for left <= right {
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		if leftMax < rightMax {
			sum += leftMax - height[left]
			left++
		} else {
			sum += rightMax - height[right]
			right--
		}
	}
	return sum
}

//接雨水  water[i]=min( max(height[0...i],  max(height[i...end])) -height[i]
func trap1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	sum := 0
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		if leftMax[i-1] < height[i] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	for j := len(height) - 2; j >= 0; j-- {
		if rightMax[j+1] < height[j] {
			rightMax[j] = height[j]
		} else {
			rightMax[j] = rightMax[j+1]
		}
	}

	for k := 1; k < len(height); k++ {
		if leftMax[k] > rightMax[k] {
			sum += rightMax[k] - height[k]
		} else {
			sum += leftMax[k] - height[k]
		}
	}
	return sum
}

//暴力方法
func trap2(height []int) int {
	sum := 0
	for i := 0; i < len(height); i++ {
		leftMax := 0
		rightMax := 0
		for j := i; j >= 0; j-- {
			if height[j] > leftMax {
				leftMax = height[j]
			}
		}
		for k := i; k < len(height); k++ {
			if height[k] > rightMax {
				rightMax = height[k]
			}
		}
		if leftMax > rightMax {
			sum += rightMax - height[i]
		} else {
			sum += leftMax - height[i]
		}
	}
	return sum
}

func Test_42(t *testing.T) {
	t.Log(trap([]int{}))                                   //0
	t.Log(trap([]int{0, 1}))                               //6
	t.Log(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) //6
}
