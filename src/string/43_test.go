package string

import "strconv"

//字符串相乘
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	//用数组存储每一位相乘的结果
	resultArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j]) - '0'
			resultArr[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		resultArr[i-1] += resultArr[i] / 10
		resultArr[i] %= 10
	}
	result := ""
	index := 0
	if resultArr[0] == 0 {
		index = 1
	}
	for ; index < m+n; index++ {
		result += strconv.Itoa(resultArr[index])
	}
	return result
}
