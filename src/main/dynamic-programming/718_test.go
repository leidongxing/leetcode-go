package dynamic_programming

import "testing"

//最长重复子数组
func findLength(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	//A 0到i  B 0到j  的重复子数组
	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, len(B)+1)
	}
	result := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				if i-1 < 0 || j-1 < 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else {
				dp[i][j] = 0
			}

			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result
}

func Test_718(t *testing.T) {
	t.Log(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7})) // [3, 2, 1]  3
}
