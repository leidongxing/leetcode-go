package recursion

import (
	"testing"
)

//把数字翻译成字符串
func translateNum(num int) int {
	//小于10  肯定只能一个一个解释
	if num < 10 {
		return 1
	}
	//末位形如 01  33 只能读一位
	if num%100 < 10 || num%100 > 25 {
		return translateNum(num / 10)
	}
	return translateNum(num/10) + translateNum(num/100)
}

func Test_46(t *testing.T) {
	t.Log(translateNum(12258)) //5
	t.Log(translateNum(26))    //1
}
