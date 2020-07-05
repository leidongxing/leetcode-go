package dynamic_programming

import (
	"sort"
	"testing"
)

//俄罗斯套娃信封问题
//二维数组下的 最长递增子序列
func maxEnvelopes1(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		//宽相同 对高进行排序
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	//保存当前最长的信封序列
	max := 1
	dp := make([]int, len(envelopes))
	dp[0] = 1
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
				if dp[i] > max {
					max = dp[i]
					break
				}
			}
		}
	}
	return max
}

//俄罗斯套娃信封问题
//二维数组下的 最长递增子序列  对顺序查找进行优化
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, 0)
	//只判断高度
	for i := 0; i < len(envelopes); i++ {
		low, high := 0, len(dp)-1
		for low <= high {
			mid := low + (high-low)/2
			if envelopes[i][1] > dp[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if low >= len(dp) {
			dp = append(dp, envelopes[i][1])
		} else {
			dp[low] = envelopes[i][1]
		}
	}
	return len(dp)
}

func Test_354(t *testing.T) {
	//t.Log(maxEnvelopes([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})) //4
	//t.Log(maxEnvelopes([][]int{{7, 8}, {1, 2}, {5, 6}, {3, 4}})) //4
	t.Log(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}})) //3
}
