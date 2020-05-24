package dp

import "testing"

func divisorGame(N int) bool {
	if N == 1 {
		return false
	}
	dp := make([]bool, N+1, N+1)
	dp[0] = false
	dp[1] = false
	dp[2] = true

	for i := 3; i < N+1; i++ {
		for j := 1; j < i; j++ {
			//如果j是i的约数 并且dp[i-j]==false
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[N]
}

//最终结果应该是占到2的赢，占到1的输；
//若当前为奇数，奇数的约数只能是奇数或者 1，因此下一个一定是偶数；
//若当前为偶数， 偶数的约数可以是奇数可以是偶数也可以是 1，因此直接减 1，则下一个是奇数；  奇则输，偶则赢。
func divisorGame1(N int) bool {
	return N%2 == 0
}

func Test_1025(t *testing.T) {
	//t.Log(divisorGame(2)) //true
	//t.Log(divisorGame(3)) //false
	t.Log(divisorGame(4)) //true
}
