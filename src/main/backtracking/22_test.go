package backtracking

import "testing"

//括号生成
func generateParenthesis(n int) []string {
	result := new([]string)
	backtracking(n, n, "", result)
	return *result
}

func backtracking(left, right int, tmp string, result *[]string) {
	//回溯跳出条件  右括号=0
	if right == 0 {
		*result = append(*result, tmp)
		return
	}
	//先生成左括号
	if left > 0 {
		backtracking(left-1, right, tmp+"(", result)
	}
	//括号都是成对出现  右括号数量大于左括号
	if right > left {
		backtracking(left, right-1, tmp+")", result)
	}
}

func Test_22(t *testing.T) {
	t.Log(generateParenthesis(3))
}
