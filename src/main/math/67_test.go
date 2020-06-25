package math

import (
	"strconv"
	"testing"
)

//二进制求和
func addBinary(a string, b string) string {
	result := ""
	nextAdd := 0
	lenA, lenB := len(a), len(b)
	max := 0
	if lenA > lenB {
		max = lenA
	} else {
		max = lenB
	}
	for i := 0; i < max; i++ {
		if i < lenA {
			nextAdd += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			nextAdd += int(b[lenB-i-1] - '0')
		}
		//从后往前插入
		result = strconv.Itoa(nextAdd%2) + result
		nextAdd /= 2
	}
	//最后还有 最多往前进一位
	if nextAdd > 0 {
		result = "1" + result
	}
	return result
}

func Test_67(t *testing.T) {
	t.Log(addBinary("11", "1"))      //100
	t.Log(addBinary("1010", "1011")) // 10101
}
