package _math

//回文数
func isPalindrome(x int) bool {
	//负数不是回文数
	if x < 0 {
		return false
	}
	old := x
	r := 0
	for x != 0 {
		r = r*10 + x%10
		x /= 10
	}
	if r == old {
		return true
	}
	return false
}
