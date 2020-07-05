package array

import "testing"

//最佳观光组合
//A[i]+ A[j]+i-j最大值  J点时 A[j]-j不会变化
func maxScoreSightseeingPair(A []int) int {
	if len(A) < 2 {
		return 0
	}

	max := 0
	maxI := A[0] + 0
	for j := 1; j < len(A); j++ {
		//更新
		if maxI+A[j]-j > max {
			max = maxI + A[j] - j
		}
		i := j
		if A[i]+i > maxI {
			maxI = A[i] + i
		}
	}
	return max
}

func maxScoreSightseeingPair2(A []int) int {
	if len(A) < 2 {
		return 0
	}
	max := 0
	dp := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		dp[i] = make([]int, len(A))
		dp[i][i] = 0
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			sum := A[i] + A[j] + i - j
			if sum > dp[i][j-1] {
				dp[i][j] = sum
			} else {
				dp[i][j] = dp[i][j-1]
			}

			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max
}

//暴力破解
func maxScoreSightseeingPair1(A []int) int {
	if len(A) < 2 {
		return 0
	}

	max := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			sum := A[i] + A[j] + i - j
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func Test_1014(t *testing.T) {
	t.Log(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})) //11
	t.Log(maxScoreSightseeingPair([]int{1, 3, 5}))       //7
}
