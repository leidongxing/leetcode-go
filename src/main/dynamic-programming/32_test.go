package dynamic_programming

import "testing"

//最长有效括号
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	result := 0
	dp := make([]int, len(s))
	//从1开始遍历
	for i := 1; i < len(s); i++ {
		//有效的子串以）结尾
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					//有效长度+2
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				//()的上一个为（
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			//更新
			result = max(result, dp[i])
		}
	}
	return result
}

func Test_32(t *testing.T) {
	t.Log(longestValidParentheses("(()"))    //2
	t.Log(longestValidParentheses(")()())")) //4
}
