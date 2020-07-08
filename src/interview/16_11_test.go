package interview

import "testing"

//跳水板
func divingBoard(shorter int, longer int, k int) []int {
	var result []int
	if k == 0 {
		return result
	}
	if shorter == longer {
		result = append(result, k*shorter)
		return result
	}
	index := 0
	for i := 0; i <= k; i++ {
		next := (k-i)*shorter + i*longer
		if index == 0 {
			result = append(result, next)
			index++
		} else if next != result[index-1] {
			result = append(result, next)
			index++
		}
	}
	return result
}

func Test_16_11(t *testing.T) {
	t.Log(divingBoard(1, 2, 3))
	t.Log(divingBoard(1, 2, 4))
}
