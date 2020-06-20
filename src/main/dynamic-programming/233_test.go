package dynamic_programming

import (
	"math"
	"strconv"
	"testing"
)

//数字1的个数
//个位上的1 十位上的1  百位上的1  千位上的1
func countDigitOne(n int) int {
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	if n <= 0 {
		return 0
	}
	s := strconv.Itoa(n)
	length := len(s)
	dp := make([]int, length+1)
	//从低位到高位
	for i := length - 1; i >= 0; i-- {
		//当前位数是1
		if s[i] == '1' {
			dp[i] = (length-i-1)*int(math.Pow10(length-i-2)) + n%int(math.Pow10(length-i-1)) + dp[i+1] + 1
		} else {
			dp[i] = (int(s[i])-'0')*(length-i-1)*int(math.Pow10(length-i-2)) + min(int(s[i])-'0', 1)*int(math.Pow10(length-i-1)) + dp[i+1]
		}
	}
	return dp[0]
}

//暴力方法 超时
func countDigitOne1(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		for j := 0; j < len(s); j++ {
			if s[j] == '1' {
				sum++
			}
		}
	}
	return sum
}

func Test_233(t *testing.T) {
	t.Log(countDigitOne(13))        //6
	t.Log(countDigitOne(824883294)) //767944060
}
