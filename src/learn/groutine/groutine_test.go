package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//复制了i的值
		go func(i int) {
			fmt.Println(i)
		}(i)

		//错误写法 共享i
		//go func(){
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Microsecond * 50)
}
