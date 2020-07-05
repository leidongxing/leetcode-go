package dynamic_programming

import "testing"

//新21点  不超过N的概率
//能抽到的最大牌点数 K-1+W
func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W)
	sum := 0.0
	//[K,K-1+W]不能抽牌  小于等于N 概率=1  大于N概率=0
	for i := K; i < K+W; i++ {
		if i <= N {
			dp[i] = 1.0
			sum += 1.0
		} else {
			dp[i] = 0.0
		}
	}
	//[0,K-1]可以抽牌
	//当前有K-1张牌 大于N的概率 等于 dp[K]+...dp[K-1+W] /W
	//当前有K-1张牌 大于N的概率 等于 dp[K+1]+dp[K]+...dp[K-2+W] /W
	for i := K - 1; i > -1; i-- {
		dp[i] = sum / (float64(W))
		//更新sum
		sum = sum - dp[i+W] + dp[i]
	}
	//没有牌时的概率
	return dp[0]
}
func Test_837(t *testing.T) {
	t.Log(new21Game(10, 1, 10))
	t.Log(new21Game(6, 1, 10))
	t.Log(new21Game(21, 17, 10))
}
