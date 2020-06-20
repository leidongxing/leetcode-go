package dynamic_programming

import "testing"

//正则表达式匹配  '.' 匹配任意单个字符  '*' 匹配零个或多个前面的那一个元素
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	//内部函数 匹配方法
	matches := func(i, j int) bool {
		//字符串是空 不匹配任何字符 直接返回false
		if i == 0 {
			return false
		}
		//模式为. 匹配多个任意字符 直接返回true
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	//使用二维数组 存储 s p长度的匹配情况
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				//需要匹配p[j-1]的前一个元素 p[j-2]
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func Test_10(t *testing.T) {
	t.Log(isMatch("aa", "a"))                   //false
	t.Log(isMatch("aa", "a*"))                  //true
	t.Log(isMatch("ab", ".*"))                  //true
	t.Log(isMatch("aab", "c*a*b"))              //true
	t.Log(isMatch("mississippi", "mis*is*p*.")) //false

}
