package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//函数可以返回多个返回值
func returnMultiValues() (int, int) {
	return rand.Intn(20), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

//可变长参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

//延迟执行  实现类似于try catch finally释放资源的效果
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	//退出程序
	panic("error")
	fmt.Println("end")
}

func Clear() {
	fmt.Println("Clear resources.")
}
