package bit

import "testing"

// 比特位计数
// 奇数的1的个数 等于 偶数的1的个数+1
// 偶数的1的个数 等于 偶数/2的1的个数  偶数/2=偶数右移1位 补0  >>1
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		//如果是最低位是0 偶数
		if i&1 == 0 {
			dp[i] = dp[i>>1]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}

func Test_338(t *testing.T) {
	t.Log(countBits(2)) //[0,1,1]
	t.Log(countBits(5)) //[0,1,1,2,1,2]
}
