package dynamic_programming

import "testing"

func Test_5(t *testing.T) {
	t.Log(longestPalindrome("babad")) //bab
	t.Log(longestPalindrome("cbbd"))  //bb
}

//动态规划 对于一个回文子串 如果它的最左边/最右边相等 会成为一个更长的回文子串
func longestPalindrome(s string) string {
	n := len(s)
	result := ""
	//二维数组初始化  dynamic-programming[i][j]表示s[i..j]是否为回文子串
	//状态转移方程   (s[i]==s[j] && dynamic-programming[i+1][j-1]==1) -----> dynamic-programming[i][j]==1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for low := 0; low < n; low++ {
		for i := 0; i+low < n; i++ {
			j := i + low
			if low == 0 {
				dp[i][j] = 1
			} else if low == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			//更新最长子串
			if dp[i][j] > 0 && low+1 > len(result) {
				result = s[i : i+low+1]
			}
		}
	}
	return result
}

//中心扩散方法 分奇数/偶数两种情况讨论
func longestPalindrome1(s string) string {
	var length = len(s)
	if length < 2 {
		return s
	}
	var index int
	var size int
	for i := 0; i < length-1; i++ {
		var j int
		var k int
		// XXXaXXX
		j, k = extendPalindrome(s, i, i)
		if size < k-j-1 {
			index = j + 1
			size = k - j - 1
		}
		// XXXaaXXX
		j, k = extendPalindrome(s, i, i+1)
		if size < k-j-1 {
			index = j + 1
			size = k - j - 1
		}

	}
	return s[index : index+size]
}

func extendPalindrome(s string, j int, k int) (int, int) {
	for j >= 0 && k < len(s) && s[j] == s[k] {
		j--
		k++
	}
	return j, k
}
