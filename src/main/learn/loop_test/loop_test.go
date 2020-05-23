package loop_test

import "testing"

//使用for循环构建while循环
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
