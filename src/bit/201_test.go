package bit

import "testing"

//数字范围按位与
func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n &= n - 1
	}
	return n
}

func Test_201(t *testing.T) {
	t.Log(rangeBitwiseAnd(5, 7)) // 101  111  101
	t.Log(rangeBitwiseAnd(0, 1))
}
