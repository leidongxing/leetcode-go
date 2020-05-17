package test

import "testing"

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1
	a := 1
	b := 1
	t.Log(a, " ")
	for i := 0; i < 5; i++ {
		t.Log("", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

//go中交换变量
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
