package _interface

import (
	"fmt"
	"testing"
)

//空接口 类似于void*
func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyIntegerAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
