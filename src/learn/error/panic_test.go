package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	fmt.Println("start")
	//退出前会进行调用
	//defer func() {
	//	fmt.Print("finally")
	//}()
	defer func() {
		//recover异常  错误恢复 继续处理
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	//会打印出调用在栈信息
	panic(errors.New("something wrong!"))
}
