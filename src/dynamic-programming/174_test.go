package dynamic_programming

import "testing"

//地下城游戏
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, len(dungeon))
	for i := 0; i < len(dungeon); i++ {
		dp[i] = make([]int, len(dungeon[0]))
	}

	if dungeon[len(dungeon)-1][len(dungeon[0])-1] >= 0 {
		dp[len(dungeon)-1][len(dungeon[0])-1] = 1
	} else {
		dp[len(dungeon)-1][len(dungeon[0])-1] = 1 - dungeon[len(dungeon)-1][len(dungeon[0])-1]
	}

	for i := len(dungeon) - 1; i >= 0; i-- {
		for j := len(dungeon[0]) - 1; j >= 0; j-- {
			if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
				continue
			} else if i == len(dungeon)-1 {
				dp[i][j] = dp[i][j+1]
				if dungeon[i][j] >= 0 {
					dp[i][j] = dp[i][j+1] - dungeon[i][j]
					if dp[i][j] <= 0 {
						dp[i][j] = 1
					}
				} else {
					dp[i][j] = dp[i][j+1] - dungeon[i][j]
				}
			} else if j == len(dungeon[0])-1 {
				dp[i][j] = dp[i+1][j]
				if dungeon[i][j] >= 0 {
					dp[i][j] = dp[i+1][j] - dungeon[i][j]
					if dp[i][j] <= 0 {
						dp[i][j] = 1
					}
				} else {
					dp[i][j] = dp[i+1][j] - dungeon[i][j]
				}
			} else {
				dp[i][j] = min(dp[i][j+1], dp[i+1][j])
				if dungeon[i][j] >= 0 {
					dp[i][j] = dp[i][j] - dungeon[i][j]
					if dp[i][j] <= 0 {
						dp[i][j] = 1
					}
				} else {
					dp[i][j] = dp[i][j] - dungeon[i][j]
				}
			}
		}
	}
	return dp[0][0]
}

func Test_174(t *testing.T) {
	t.Log(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}})) //7
}
