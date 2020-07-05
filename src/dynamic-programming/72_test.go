package dynamic_programming

import (
	"math"
	"testing"
)

//编辑距离
//插入 删除 替换一个字符  word1 --> word2
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	//有空字符 直接插入即为最小距离
	if m == 0 || n == 0 {
		return n + m
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//word1 删除一个元素  距离+1
			delete := dp[i-1][j] + 1
			//word1 插入一个元素  距离+1
			insert := dp[i][j-1] + 1
			//word1 替换元素 距离+1
			replace := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				replace += 1
			}
			dp[i][j] = int(math.Min(float64(delete), math.Min(float64(insert), float64(replace))))
		}
	}
	return dp[m][n]
}

func Test_72(t *testing.T) {
	t.Log(minDistance("horse", "ros"))           //3  horse->rorse->rose->ros
	t.Log(minDistance("intention", "execution")) //5  intention->inention->exention->exection->execution
}
