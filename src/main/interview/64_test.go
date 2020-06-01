package interview

import "testing"

func sumNums(n int) int {
	if n <= 0 {
		return n
	}
	return sumNums(n-1) + n
}

func Test_64(t *testing.T) {
	t.Log(sumNums(3))
	t.Log(sumNums(9))
}
